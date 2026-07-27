package shift

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/rajeshbond/smart/internal/http/device/device_data/shiftprovider"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Create(
	ctx context.Context,
	tx *sql.Tx,
	userID int64,
	req CreateShiftTimingRequest,
) error {

	for _, t := range req.Timings {
		if err := s.InsertShiftTiming(ctx, tx, req.TenantShiftID, t, userID); err != nil {
			return err
		}
	}

	return nil
}

func (s *Store) InsertShiftTiming(
	ctx context.Context,
	tx *sql.Tx,
	shiftID int64,
	t TimingDTO,
	userID int64,
) error {

	query := `
	INSERT INTO shift_timing
	(
		tenant_shift_id,
		shift_start,
		shift_end,
		weekday,
		created_by,
		updated_by
	)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (tenant_shift_id, weekday, shift_start, shift_end)
	DO NOTHING;
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		shiftID,
		t.ShiftStart,
		t.ShiftEnd,
		t.Weekday,
		userID,
		userID,
	)

	if err != nil {
		// 🔥 Handle DB trigger error nicely
		if pqErr, ok := err.(*pq.Error); ok {
			return fmt.Errorf("shift timing error: %s", pqErr.Message)
		}
		return err
	}

	return nil
}

func (s *Store) GetTenantIDByCode(ctx context.Context, tx *sql.Tx, code string) (int64, error) {

	var id int64

	err := tx.QueryRowContext(ctx,
		`SELECT id FROM tenant WHERE tenant_code = $1`,
		code,
	).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("tenant not found: %s", code)
	}

	return id, err
}

func (s *Store) UpsertTenantShift(ctx context.Context, tx *sql.Tx, tenantID int64, shiftName string, userID int64) (int64, error) {

	query := `
	INSERT INTO tenant_shift (tenant_id, shift_name, created_by, updated_by)
	VALUES ($1,$2,$3,$4)
	ON CONFLICT (tenant_id, shift_name)
	DO UPDATE SET updated_at = NOW()
	RETURNING id;
	`

	var id int64
	err := tx.QueryRowContext(ctx, query,
		tenantID,
		shiftName,
		userID,
		userID,
	).Scan(&id)

	return id, err
}

func (s *Store) GetProductionDay(
	ctx context.Context,
	tenantID int64,
	productionTime time.Time,
) (*shiftprovider.ProductionDayInfo, error) {

	rows, err := s.db.QueryContext(
		ctx,
		GetShiftByTenant,
		tenantID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	type shiftRow struct {
		ShiftID   int64
		ShiftName string
		StartTime time.Time
		EndTime   time.Time
		Weekday   int
	}

	var shifts []shiftRow

	for rows.Next() {

		var item shiftRow

		err = rows.Scan(
			&item.ShiftID,
			&item.ShiftName,
			&item.StartTime,
			&item.EndTime,
			&item.Weekday,
		)

		if err != nil {
			return nil, err
		}

		shifts = append(shifts, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(shifts) == 0 {
		return nil, sql.ErrNoRows
	}

	//----------------------------------------------------------
	// Convert UTC -> IST
	//----------------------------------------------------------

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return nil, err
	}

	localProductionTime := productionTime.In(loc)

	//----------------------------------------------------------
	// Find earliest shift start
	//----------------------------------------------------------

	var earliest time.Time
	first := true

	for _, sh := range shifts {

		if first || sh.StartTime.Before(earliest) {

			earliest = sh.StartTime
			first = false
		}
	}

	//----------------------------------------------------------
	// Build today's production day start
	//----------------------------------------------------------

	dayStart := time.Date(
		localProductionTime.Year(),
		localProductionTime.Month(),
		localProductionTime.Day(),
		earliest.Hour(),
		earliest.Minute(),
		0,
		0,
		loc,
	)

	//----------------------------------------------------------
	// Before production day starts?
	//----------------------------------------------------------

	if localProductionTime.Before(dayStart) {
		dayStart = dayStart.AddDate(0, 0, -1)
	}

	dayEnd := dayStart.Add(24 * time.Hour)

	return &shiftprovider.ProductionDayInfo{
		DayStart: dayStart,
		DayEnd:   dayEnd,
	}, nil
}

func (s *Store) GetShiftByProductionTime(
	ctx context.Context,
	tenantID int64,
	productionTime time.Time,
) (*shiftprovider.ShiftInfo, error) {

	rows, err := s.db.QueryContext(
		ctx,
		GetShiftByTenant,
		tenantID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type shiftRow struct {
		ShiftID   int64
		ShiftName string
		StartTime time.Time
		EndTime   time.Time
		Weekday   int
	}

	var shifts []shiftRow

	for rows.Next() {

		var item shiftRow

		err = rows.Scan(
			&item.ShiftID,
			&item.ShiftName,
			&item.StartTime,
			&item.EndTime,
			&item.Weekday,
		)

		if err != nil {
			return nil, err
		}

		shifts = append(shifts, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(shifts) == 0 {
		return nil, sql.ErrNoRows
	}

	//---------------------------------------------------------
	// Go Sunday=0
	// PostgreSQL Monday=1...Sunday=7
	//---------------------------------------------------------

	weekday := int(productionTime.Weekday())

	if weekday == 0 {
		weekday = 7
	}

	for _, sh := range shifts {

		if sh.Weekday != weekday {
			continue
		}

		start := time.Date(
			productionTime.Year(),
			productionTime.Month(),
			productionTime.Day(),
			sh.StartTime.Hour(),
			sh.StartTime.Minute(),
			0,
			0,
			productionTime.Location(),
		)

		end := time.Date(
			productionTime.Year(),
			productionTime.Month(),
			productionTime.Day(),
			sh.EndTime.Hour(),
			sh.EndTime.Minute(),
			0,
			0,
			productionTime.Location(),
		)

		//-------------------------------------------------
		// Overnight Shift
		//-------------------------------------------------

		if !end.After(start) {

			if productionTime.Before(end) {

				start = start.AddDate(0, 0, -1)

			} else {

				end = end.AddDate(0, 0, 1)
			}
		}

		if productionTime.Equal(start) ||
			(productionTime.After(start) &&
				productionTime.Before(end)) {

			// log.Println("Shift Found")
			// log.Println("Shift ID ----->", sh.ShiftID)
			// log.Println("Shift Name ----->", sh.ShiftName)

			return &shiftprovider.ShiftInfo{

				ShiftID: sh.ShiftID,

				ShiftName: sh.ShiftName,

				ShiftStart: start,

				ShiftEnd: end,
			}, nil
		}
	}

	return nil, fmt.Errorf("shift not found")
}

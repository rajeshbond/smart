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

	return s.InsertShiftTiming(
		ctx,
		tx,
		req.TenantShiftID,
		req.ShiftStart,
		req.ShiftEnd,
		userID,
	)
}

func (s *Store) InsertShiftTiming(
	ctx context.Context,
	tx *sql.Tx,
	shiftID int64,
	shiftStart string,
	shiftEnd string,
	userID int64,
) error {

	query := `
	INSERT INTO shift_timing
	(
		tenant_shift_id,
		shift_start,
		shift_end,
		created_by,
		updated_by
	)
	VALUES
	(
		$1, $2, $3, $4, $5
	)
	ON CONFLICT (tenant_shift_id)
	DO UPDATE
	SET
		shift_start = EXCLUDED.shift_start,
		shift_end   = EXCLUDED.shift_end,
		updated_by  = EXCLUDED.updated_by,
		updated_at  = NOW();
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		shiftID,
		shiftStart,
		shiftEnd,
		userID,
		userID,
	)

	if err != nil {
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
		ShiftID    int64
		ShiftName  string
		ShiftStart string
		ShiftEnd   string
	}

	var shifts []shiftRow

	for rows.Next() {

		var item shiftRow

		err := rows.Scan(
			&item.ShiftID,
			&item.ShiftName,
			&item.ShiftStart,
			&item.ShiftEnd,
		)
		if err != nil {
			return nil, err
		}

		shifts = append(shifts, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(shifts) == 0 {
		return nil, sql.ErrNoRows
	}

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return nil, err
	}

	localTime := productionTime.In(loc)

	//--------------------------------------------------
	// Find earliest shift start
	//--------------------------------------------------

	earliestMin := 24 * 60

	for _, sh := range shifts {

		min, err := toMinutes(sh.ShiftStart)
		if err != nil {
			return nil, err
		}

		if min < earliestMin {
			earliestMin = min
		}
	}

	dayStart := time.Date(
		localTime.Year(),
		localTime.Month(),
		localTime.Day(),
		earliestMin/60,
		earliestMin%60,
		0,
		0,
		loc,
	)

	// If current time is before production day start,
	// production day actually started yesterday.
	if localTime.Before(dayStart) {
		dayStart = dayStart.AddDate(0, 0, -1)
	}

	dayEnd := dayStart.Add(24 * time.Hour)

	return &shiftprovider.ProductionDayInfo{
		DayStart: dayStart,
		DayEnd:   dayEnd,
	}, nil
}

// func (s *Store) GetProductionDay(ctx context.Context, tenantID int64, productionTime time.Time) (*shiftprovider.ProductionDayInfo, error) {

// 	rows, err := s.db.QueryContext(
// 		ctx,
// 		GetShiftByTenant,
// 		tenantID,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	type shiftRow struct {
// 		ShiftID   int64
// 		ShiftName string
// 		StartTime time.Time
// 		EndTime   time.Time
// 		Weekday   int
// 	}

// 	var shifts []shiftRow

// 	for rows.Next() {

// 		var item shiftRow

// 		err = rows.Scan(
// 			&item.ShiftID,
// 			&item.ShiftName,
// 			&item.StartTime,
// 			&item.EndTime,
// 			&item.Weekday,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		shifts = append(shifts, item)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	if len(shifts) == 0 {
// 		return nil, sql.ErrNoRows
// 	}

// 	//----------------------------------------------------------
// 	// Convert UTC -> IST
// 	//----------------------------------------------------------

// 	loc, err := time.LoadLocation("Asia/Kolkata")
// 	if err != nil {
// 		return nil, err
// 	}

// 	localProductionTime := productionTime.In(loc)

// 	//----------------------------------------------------------
// 	// Find earliest shift start
// 	//----------------------------------------------------------

// 	var earliest time.Time
// 	first := true

// 	for _, sh := range shifts {

// 		if first || sh.StartTime.Before(earliest) {

// 			earliest = sh.StartTime
// 			first = false
// 		}
// 	}

// 	//----------------------------------------------------------
// 	// Build today's production day start
// 	//----------------------------------------------------------

// 	dayStart := time.Date(
// 		localProductionTime.Year(),
// 		localProductionTime.Month(),
// 		localProductionTime.Day(),
// 		earliest.Hour(),
// 		earliest.Minute(),
// 		0,
// 		0,
// 		loc,
// 	)

// 	//----------------------------------------------------------
// 	// Before production day starts?
// 	//----------------------------------------------------------

// 	if localProductionTime.Before(dayStart) {
// 		dayStart = dayStart.AddDate(0, 0, -1)
// 	}

// 	dayEnd := dayStart.Add(24 * time.Hour)

// 	return &shiftprovider.ProductionDayInfo{
// 		DayStart: dayStart,
// 		DayEnd:   dayEnd,
// 	}, nil

// }

//

//

// func (s *Store) GetShiftByProductionTime(ctx context.Context, tenantID int64, productionTime time.Time) (*shiftprovider.ShiftInfo, error) {

// 	rows, err := s.db.QueryContext(
// 		ctx,
// 		GetShiftByTenant,
// 		tenantID,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	type shiftRow struct {
// 		ShiftID   int64
// 		ShiftName string
// 		StartTime time.Time
// 		EndTime   time.Time
// 		Weekday   int
// 	}

// 	var shifts []shiftRow

// 	for rows.Next() {

// 		var item shiftRow

// 		err = rows.Scan(
// 			&item.ShiftID,
// 			&item.ShiftName,
// 			&item.StartTime,
// 			&item.EndTime,
// 			&item.Weekday,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		shifts = append(shifts, item)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	if len(shifts) == 0 {
// 		return nil, sql.ErrNoRows
// 	}

// 	//---------------------------------------------------------
// 	// Go Sunday=0
// 	// PostgreSQL Monday=1...Sunday=7
// 	//---------------------------------------------------------

// 	weekday := int(productionTime.Weekday())

// 	if weekday == 0 {
// 		weekday = 7
// 	}

// 	for _, sh := range shifts {

// 		if sh.Weekday != weekday {
// 			continue
// 		}

// 		start := time.Date(
// 			productionTime.Year(),
// 			productionTime.Month(),
// 			productionTime.Day(),
// 			sh.StartTime.Hour(),
// 			sh.StartTime.Minute(),
// 			0,
// 			0,
// 			productionTime.Location(),
// 		)

// 		end := time.Date(
// 			productionTime.Year(),
// 			productionTime.Month(),
// 			productionTime.Day(),
// 			sh.EndTime.Hour(),
// 			sh.EndTime.Minute(),
// 			0,
// 			0,
// 			productionTime.Location(),
// 		)

// 		//-------------------------------------------------
// 		// Overnight Shift
// 		//-------------------------------------------------

// 		if !end.After(start) {

// 			if productionTime.Before(end) {

// 				start = start.AddDate(0, 0, -1)

// 			} else {

// 				end = end.AddDate(0, 0, 1)
// 			}
// 		}

// 		if productionTime.Equal(start) ||
// 			(productionTime.After(start) &&
// 				productionTime.Before(end)) {

// 			// log.Println("Shift Found")
// 			// log.Println("Shift ID ----->", sh.ShiftID)
// 			// log.Println("Shift Name ----->", sh.ShiftName)

// 			return &shiftprovider.ShiftInfo{

// 				ShiftID: sh.ShiftID,

// 				ShiftName: sh.ShiftName,

// 				ShiftStart: start,

// 				ShiftEnd: end,
// 			}, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("shift not found")
// }

// func (s *Store) GetShiftByProductionTime1(
// 	ctx context.Context,
// 	tenantID int64,
// 	productionTime time.Time,
// ) (*shiftprovider.ShiftInfo, error) {

// 	rows, err := s.db.QueryContext(
// 		ctx,
// 		GetShiftByTenant,
// 		tenantID,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	type shiftRow struct {
// 		ShiftID   int64
// 		ShiftName string
// 		StartTime time.Time
// 		EndTime   time.Time
// 		Weekday   int
// 	}

// 	var shifts []shiftRow

// 	for rows.Next() {

// 		var item shiftRow

// 		err = rows.Scan(
// 			&item.ShiftID,
// 			&item.ShiftName,
// 			&item.StartTime,
// 			&item.EndTime,
// 			&item.Weekday,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		shifts = append(shifts, item)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	if len(shifts) == 0 {
// 		return nil, sql.ErrNoRows
// 	}

// 	//---------------------------------------------------------
// 	// Convert UTC -> IST
// 	//---------------------------------------------------------

// 	loc, err := time.LoadLocation("Asia/Kolkata")
// 	if err != nil {
// 		return nil, err
// 	}

// 	localProductionTime := productionTime.In(loc)

// 	//---------------------------------------------------------
// 	// Go Sunday=0
// 	// PostgreSQL Monday=1 ... Sunday=7
// 	//---------------------------------------------------------

// 	weekday := int(localProductionTime.Weekday())

// 	if weekday == 0 {
// 		weekday = 7
// 	}

// 	for _, sh := range shifts {

// 		if sh.Weekday != weekday {
// 			continue
// 		}

// 		start := time.Date(
// 			localProductionTime.Year(),
// 			localProductionTime.Month(),
// 			localProductionTime.Day(),
// 			sh.StartTime.Hour(),
// 			sh.StartTime.Minute(),
// 			0,
// 			0,
// 			loc,
// 		)

// 		end := time.Date(
// 			localProductionTime.Year(),
// 			localProductionTime.Month(),
// 			localProductionTime.Day(),
// 			sh.EndTime.Hour(),
// 			sh.EndTime.Minute(),
// 			0,
// 			0,
// 			loc,
// 		)

// 		//-------------------------------------------------
// 		// Overnight Shift
// 		//-------------------------------------------------

// 		if !end.After(start) {

// 			if localProductionTime.Before(end) {

// 				start = start.AddDate(0, 0, -1)

// 			} else {

// 				end = end.AddDate(0, 0, 1)
// 			}
// 		}

// 		if localProductionTime.Equal(start) ||
// 			(localProductionTime.After(start) &&
// 				localProductionTime.Before(end)) {

// 			return &shiftprovider.ShiftInfo{
// 				ShiftID:    sh.ShiftID,
// 				ShiftName:  sh.ShiftName,
// 				ShiftStart: start,
// 				ShiftEnd:   end,
// 			}, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("shift not found")
// }

func (s *Store) GetTenantShifts(
	ctx context.Context,
	tx *sql.Tx,
	tenantID int64,
) ([]ShiftInfo, error) {

	// query := `
	// SELECT
	// 	ts.id,
	// 	ts.shift_name,
	// 	TO_CHAR(st.shift_start, 'HH24:MI') AS shift_start,
	// 	TO_CHAR(st.shift_end, 'HH24:MI') AS shift_end
	// FROM tenant_shift ts
	// INNER JOIN shift_timing st
	// 	ON st.tenant_shift_id = ts.id
	// WHERE ts.tenant_id = $1
	// ORDER BY st.shift_start;
	// `

	rows, err := tx.QueryContext(
		ctx,
		GetShiftByTenant,
		tenantID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []ShiftInfo

	for rows.Next() {

		var item ShiftInfo

		if err := rows.Scan(
			&item.ShiftID,
			&item.ShiftName,
			&item.ShiftStart,
			&item.ShiftEnd,
		); err != nil {
			return nil, err
		}

		shifts = append(shifts, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shifts, nil
}

// ----->
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
		ShiftID    int64
		ShiftName  string
		ShiftStart string
		ShiftEnd   string
	}

	var shifts []shiftRow

	for rows.Next() {

		var sh shiftRow

		if err := rows.Scan(
			&sh.ShiftID,
			&sh.ShiftName,
			&sh.ShiftStart,
			&sh.ShiftEnd,
		); err != nil {
			return nil, err
		}

		shifts = append(shifts, sh)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(shifts) == 0 {
		return nil, fmt.Errorf("no shifts configured")
	}

	//---------------------------------------------------------
	// Tenant Timezone
	//---------------------------------------------------------

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return nil, err
	}

	localTime := productionTime.In(loc)

	for _, sh := range shifts {

		startMin, err := toMinutes(sh.ShiftStart)
		if err != nil {
			return nil, err
		}

		endMin, err := toMinutes(sh.ShiftEnd)
		if err != nil {
			return nil, err
		}

		start := time.Date(
			localTime.Year(),
			localTime.Month(),
			localTime.Day(),
			startMin/60,
			startMin%60,
			0,
			0,
			loc,
		)

		end := time.Date(
			localTime.Year(),
			localTime.Month(),
			localTime.Day(),
			endMin/60,
			endMin%60,
			0,
			0,
			loc,
		)

		//------------------------------------------
		// Overnight shift
		//------------------------------------------

		if endMin <= startMin {

			end = end.Add(24 * time.Hour)

			if localTime.Before(start) {
				start = start.Add(-24 * time.Hour)
				end = end.Add(-24 * time.Hour)
			}
		}

		if !localTime.Before(start) && localTime.Before(end) {

			return &shiftprovider.ShiftInfo{
				ShiftID:    sh.ShiftID,
				ShiftName:  sh.ShiftName,
				ShiftStart: start,
				ShiftEnd:   end,
			}, nil
		}
	}

	return nil, fmt.Errorf("shift not found")
}

// ---->

// func (s *Store) GetShiftByProductionTime(
// 	ctx context.Context,
// 	tenantID int64,
// 	productionTime time.Time,
// ) (*shiftprovider.ShiftInfo, error) {

// 	rows, err := s.db.QueryContext(
// 		ctx,
// 		GetShiftByTenant,
// 		tenantID,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	type shiftRow struct {
// 		ShiftID    int64
// 		ShiftName  string
// 		ShiftStart string
// 		ShiftEnd   string
// 	}

// 	var shifts []shiftRow

// 	for rows.Next() {

// 		var item shiftRow

// 		err = rows.Scan(
// 			&item.ShiftID,
// 			&item.ShiftName,
// 			&item.ShiftStart,
// 			&item.ShiftEnd,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		shifts = append(shifts, item)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	if len(shifts) == 0 {
// 		return nil, fmt.Errorf("no shifts configured")
// 	}

// 	//---------------------------------------------------------
// 	// TODO:
// 	// Later load tenant timezone from tenant.timezone
// 	//---------------------------------------------------------

// 	loc, err := time.LoadLocation("Asia/Kolkata")
// 	if err != nil {
// 		return nil, err
// 	}

// 	localTime := productionTime.In(loc)

// 	for _, sh := range shifts {

// 		startMin, err := toMinutes(sh.ShiftStart)
// 		if err != nil {
// 			return nil, err
// 		}

// 		endMin, err := toMinutes(sh.ShiftEnd)
// 		if err != nil {
// 			return nil, err
// 		}

// 		start := time.Date(
// 			localTime.Year(),
// 			localTime.Month(),
// 			localTime.Day(),
// 			startMin/60,
// 			startMin%60,
// 			0,
// 			0,
// 			loc,
// 		)

// 		end := time.Date(
// 			localTime.Year(),
// 			localTime.Month(),
// 			localTime.Day(),
// 			endMin/60,
// 			endMin%60,
// 			0,
// 			0,
// 			loc,
// 		)

// 		//-----------------------------------------
// 		// Overnight Shift
// 		//-----------------------------------------

// 		if endMin <= startMin {

// 			if localTime.Before(end) {

// 				start = start.AddDate(0, 0, -1)

// 			} else {

// 				end = end.AddDate(0, 0, 1)
// 			}
// 		}

// 		if (localTime.Equal(start) || localTime.After(start)) &&
// 			localTime.Before(end) {

// 			return &shiftprovider.ShiftInfo{
// 				ShiftID:    sh.ShiftID,
// 				ShiftName:  sh.ShiftName,
// 				ShiftStart: start,
// 				ShiftEnd:   end,
// 			}, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("shift not found")
// }

//

// func (s *Store) GetShiftByProductionTime(
// 	ctx context.Context,
// 	tenantID int64,
// 	productionTime time.Time,
// ) (*shiftprovider.ShiftInfo, error) {

// 	rows, err := s.db.QueryContext(
// 		ctx,
// 		GetShiftByTenant,
// 		tenantID,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	type shiftRow struct {
// 		ShiftID   int64
// 		ShiftName string
// 		StartTime time.Time
// 		EndTime   time.Time
// 	}

// 	var shifts []shiftRow

// 	for rows.Next() {

// 		var item shiftRow

// 		err = rows.Scan(
// 			&item.ShiftID,
// 			&item.ShiftName,
// 			&item.StartTime,
// 			&item.EndTime,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		shifts = append(shifts, item)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	if len(shifts) == 0 {
// 		return nil, sql.ErrNoRows
// 	}

// 	//---------------------------------------------------------
// 	// Load Tenant Timezone
// 	//---------------------------------------------------------

// 	loc, err := time.LoadLocation("Asia/Kolkata")
// 	if err != nil {
// 		return nil, err
// 	}

// 	localProductionTime := productionTime.In(loc)

// 	//---------------------------------------------------------
// 	// Find Matching Shift
// 	//---------------------------------------------------------

// 	for _, sh := range shifts {

// 		start := time.Date(
// 			localProductionTime.Year(),
// 			localProductionTime.Month(),
// 			localProductionTime.Day(),
// 			sh.StartTime.Hour(),
// 			sh.StartTime.Minute(),
// 			0,
// 			0,
// 			loc,
// 		)

// 		end := time.Date(
// 			localProductionTime.Year(),
// 			localProductionTime.Month(),
// 			localProductionTime.Day(),
// 			sh.EndTime.Hour(),
// 			sh.EndTime.Minute(),
// 			0,
// 			0,
// 			loc,
// 		)

// 		//-----------------------------------------
// 		// Overnight Shift (20:00 -> 08:00)
// 		//-----------------------------------------

// 		if !end.After(start) {

// 			if localProductionTime.Before(end) {
// 				start = start.AddDate(0, 0, -1)
// 			} else {
// 				end = end.AddDate(0, 0, 1)
// 			}
// 		}

// 		if !localProductionTime.Before(start) &&
// 			localProductionTime.Before(end) {

// 			return &shiftprovider.ShiftInfo{
// 				ShiftID:    sh.ShiftID,
// 				ShiftName:  sh.ShiftName,
// 				ShiftStart: start,
// 				ShiftEnd:   end,
// 			}, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("shift not found")
// }

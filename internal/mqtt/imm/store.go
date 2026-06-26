package imm

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	DB *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{DB: db}
}

// InsertTelemetry logs incoming production data with context cancellation support
func (s *Store) InsertTelemetry(ctx context.Context, t IMMTelemetry) error {
	// 🛠️ FIXED: Matched column count to argument count. (Removed raw NOW() in VALUES to avoid column mismatch)
	query := `
		INSERT INTO machine_data (tenant_id, machine_id, temperature, pressure, cycle_time, status, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, NOW())`

	// 🌟 SCALE FIX: Using ctx instead of context.Background() ensures query terminates if worker drops out
	_, err := s.DB.Exec(ctx, query,
		t.TenantID,
		t.MachineID,
		t.Temperature,
		t.Pressure,
		t.CycleTime,
		t.Status,
	)

	return err
}

// UpsertHeartbeat records alive status strings from devices
func (s *Store) UpsertHeartbeat(ctx context.Context, tenantID int64, machineID string) error {
	// 🛠️ FIXED: Corrected syntax typos (NOW without parentheses, unclosed 'ONLINE single-quotes, and 'PNLINE' typo)
	query := `
		INSERT INTO machine_status (tenant_id, machine_id, last_seen, status)
		VALUES ($1, $2, NOW(), 'ONLINE')
		ON CONFLICT (tenant_id, machine_id)
		DO UPDATE SET last_seen = NOW(), status = 'ONLINE'`

	_, err := s.DB.Exec(ctx, query, tenantID, machineID)
	return err
}

// GetMachines fetches all live machines to support your monitoring watchdog worker
func (s *Store) GetMachines(ctx context.Context) ([]struct {
	TenantID  int64
	MachineID string
	LastSeen  time.Time
}, error) {

	query := `SELECT tenant_id, machine_id, last_seen FROM machine_status`

	// Using pgxpool.Query with structural rows compilation
	rows, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []struct {
		TenantID  int64
		MachineID string
		LastSeen  time.Time
	}

	for rows.Next() {
		var r struct {
			TenantID  int64
			MachineID string
			LastSeen  time.Time
		}

		// 🛠️ FIXED: Checked error return on rows.Scan to avoid corrupted allocations under heavy pool pressure
		if err := rows.Scan(&r.TenantID, &r.MachineID, &r.LastSeen); err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	// 🛠️ FIXED: Always check rows.Err() after loops when scanning multiple database allocations
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

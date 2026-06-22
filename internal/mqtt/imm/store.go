package imm

import (
	"context"
	"time"

	"github.com/rajeshbond/smart/database"
)

type Store struct {
	DB *database.DB
}

func NewStore(db *database.DB) *Store {
	return &Store{DB: db}
}

// Save telemetry

func (s *Store) InsertTelemtry(t IMMTelemetry) error {
	query := `INSERT INTO machine_data(machine_id,temperature,pressure,cycle_time,status,created_at) VALUES($1,$2,$3,$4,$5,$6,NOW())`
	_, error := s.DB.PGX.Exec(context.Background(), query,
		t.TenantID,
		t.MachineID,
		t.Temperature,
		t.Pressure,
		t.CycleTime,
		t.Status,
	)

	return error
}

//  Heartbeat update

func (s *Store) UpsertHeartbeat(tenantID int64, machineID string) error {
	query := `
	INSERT INTO machine_status(tenant_id,machine_id,last_seen,status)VALUES($1,$2,NOW,'ONLINE)
	ON CONFLICT (tenant_id,machine_id)
	DO UPDATE SET last_seen = NOW(),status='PNLINE'`
	_, err := s.DB.PGX.Exec(context.Background(), query,
		tenantID, machineID,
	)

	return err
}

// Get machines for watchdog

func (s *Store) GetMAchines() ([]struct {
	TenantID  int64
	MachineID string
	LastSeen  time.Time
}, error) {

	query := `
SELECT tenant_id,machine_id,last_seen FROM machine_status`

	rows, err := s.DB.PGX.Query(context.Background(), query)

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

		rows.Scan(&r.TenantID, &r.MachineID, &r.LastSeen)
		result = append(result, r)
	}

	return result, nil

}

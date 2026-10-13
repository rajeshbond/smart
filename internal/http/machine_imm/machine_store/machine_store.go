package machinestore

import "database/sql"

type MachineStore struct {
	db *sql.DB
}

func NewMachineStore(db *sql.DB) *MachineStore {
	return &MachineStore{db: db}
}

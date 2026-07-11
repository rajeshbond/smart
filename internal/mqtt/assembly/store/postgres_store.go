package store

import "database/sql"

type PostgresProductionStore struct {
	db *sql.DB
}

func NewProductionStore(db *sql.DB) ProductionStore {
	return &PostgresProductionStore{
		db: db,
	}
}

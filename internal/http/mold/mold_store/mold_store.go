package moldstore

import "database/sql"

type MoldStore struct {
	db *sql.DB
}

func NewMoldStore(db *sql.DB) *MoldStore {
	return &MoldStore{db: db}
}

package store

import "database/sql"

type dept_store struct {
	db *sql.DB
}

func NewDeptStore(db *sql.DB) *dept_store {
	return &dept_store{db: db}
}

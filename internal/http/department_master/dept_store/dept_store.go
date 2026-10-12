package deptstore

import (
	"context"
	"database/sql"
)

type DeptStore struct {
	db *sql.DB
}

func NewDeptStore(db *sql.DB) *DeptStore {
	return &DeptStore{db: db}
}

func (s *DeptStore) BeginTx(
	ctx context.Context,
) (*sql.Tx, error) {

	return s.db.BeginTx(
		ctx,
		&sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		},
	)
}

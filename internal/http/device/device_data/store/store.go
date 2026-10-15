package store

import (
	"context"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) BeginTx(
	ctx context.Context,
	opts *sql.TxOptions,
) (*sql.Tx, error) {
	return s.db.BeginTx(ctx, opts)
}

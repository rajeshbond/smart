package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) BeginTx(
	ctx context.Context,
) (*sqlx.Tx, error) {

	return s.db.BeginTxx(ctx, nil)
}

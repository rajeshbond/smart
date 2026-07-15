// 1. Create

package store

import (
	"github.com/jmoiron/sqlx"
)

type PostgresStore struct {
	db *sqlx.DB
}

// Compile - time interface check 

var _ Store = (*PostgresStore)(nil)


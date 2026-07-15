// 1. Create

package store

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	devicemaster "github.com/rajeshbond/smart/internal/http/device/device_master"
	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

type PostgresStore struct {
	db *sqlx.DB
}

// Compile - time interface check

var _ Store = (*PostgresStore)(nil)

func NewPostgresStore(db *sqlx.DB) Store {
	return &PostgresStore{
		db: db,
	}
}

func (s *PostgresStore) Create(
	ctx context.Context,
	device *model.Device,
) (int64, error) {
	rows, err := s.db.NamedQueryContext(
		ctx,
		devicemaster.InsertDevice, device,
	)
	if err != nil {
		return 0, fmt.Errorf("Create device: %w", err)
	}

	defer rows.Close()

	var id int64

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("scan inserted id: %w", err)
		}
	}
	return id, nil
}

/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : create.go
 *
 * DESCRIPTION :
 * Create Device
 *
 ******************************************************************************/

package store

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

func (s *Store) Create(
	ctx context.Context,
	device *model.Device,
) (int64, error) {

	rows, err := s.db.NamedQueryContext(
		ctx,
		CreateDevice,
		device,
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int64

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

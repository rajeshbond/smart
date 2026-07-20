/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : list.go
 *
 ******************************************************************************/

package store

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// List
//------------------------------------------------------------------------------

func (s *Store) List(
	ctx context.Context,
) ([]*model.Device, error) {

	rows, err := s.db.QueryContext(
		ctx,
		ListDevices,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := make([]*model.Device, 0)

	for rows.Next() {

		device, err := scanDevice(rows)
		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}

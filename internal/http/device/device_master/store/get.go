/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : get.go
 *
 ******************************************************************************/

package store

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// Get Device By ID
//------------------------------------------------------------------------------

func (s *Store) GetByID(
	ctx context.Context,
	id int64,
) (*model.Device, error) {

	row := s.db.QueryRowContext(
		ctx,
		GetDeviceByID,
		id,
	)

	return scanDevice(row)
}

//------------------------------------------------------------------------------
// Get Device By MQTT Username
//------------------------------------------------------------------------------

func (s *Store) GetByMQTTUsername(
	ctx context.Context,
	username string,
) (*model.Device, error) {

	row := s.db.QueryRowContext(
		ctx,
		GetDeviceByMQTTUsername,
		username,
	)

	return scanDevice(row)
}

//------------------------------------------------------------------------------
// Get Device By Device Secret
//------------------------------------------------------------------------------

func (s *Store) GetBySecret(
	ctx context.Context,
	secret string,
) (*model.Device, error) {

	row := s.db.QueryRowContext(
		ctx,
		GetDeviceBySecret,
		secret,
	)

	return scanDevice(row)
}

//------------------------------------------------------------------------------
// Get Device By Chip ID
//------------------------------------------------------------------------------

func (s *Store) GetByChipID(
	ctx context.Context,
	chipID string,
) (*model.Device, error) {

	row := s.db.QueryRowContext(
		ctx,
		GetDeviceByChipID,
		chipID,
	)

	return scanDevice(row)
}

func (s *Store) GetByDeviceID(
	ctx context.Context,
	deviceID string,
) (*model.Device, error) {

	row := s.db.QueryRowContext(
		ctx,
		GetDeviceByDeviceID,
		deviceID,
	)

	return scanDevice(row)
}

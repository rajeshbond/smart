/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : get.go
 *
 * DESCRIPTION :
 * Get Device
 *
 ******************************************************************************/

package store

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

func (s *Store) GetByID(
	ctx context.Context,
	id int64,
) (*model.Device, error) {

	var device model.Device

	err := s.db.GetContext(
		ctx,
		&device,
		GetDeviceByID,
		id,
	)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (s *Store) GetByDeviceID(
	ctx context.Context,
	deviceID string,
) (*model.Device, error) {

	var device model.Device

	err := s.db.GetContext(
		ctx,
		&device,
		GetDeviceByDeviceID,
		deviceID,
	)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (s *Store) GetBySerialNumber(
	ctx context.Context,
	serialNumber string,
) (*model.Device, error) {

	var device model.Device

	err := s.db.GetContext(
		ctx,
		&device,
		GetDeviceBySerialNumber,
		serialNumber,
	)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (s *Store) GetByMQTTUsername(
	ctx context.Context,
	mqttUsername string,
) (*model.Device, error) {

	var device model.Device

	err := s.db.GetContext(
		ctx,
		&device,
		GetDeviceByMQTTUsername,
		mqttUsername,
	)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (s *Store) GetByChipID(
	ctx context.Context,
	chipID string,
) (*model.Device, error) {

	var device model.Device

	err := s.db.GetContext(
		ctx,
		&device,
		GetDeviceByChipID,
		chipID,
	)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (s *Store) ExistsByDeviceID(
	ctx context.Context,
	deviceID string,
) (bool, error) {

	var exists bool

	err := s.db.GetContext(
		ctx,
		&exists,
		ExistsByDeviceID,
		deviceID,
	)

	return exists, err
}

func (s *Store) ExistsBySerialNumber(
	ctx context.Context,
	serialNumber string,
) (bool, error) {

	var exists bool

	err := s.db.GetContext(
		ctx,
		&exists,
		ExistsBySerialNumber,
		serialNumber,
	)

	return exists, err
}

func (s *Store) ExistsByMQTTUsername(
	ctx context.Context,
	mqttUsername string,
) (bool, error) {

	var exists bool

	err := s.db.GetContext(
		ctx,
		&exists,
		ExistsByMQTTUsername,
		mqttUsername,
	)

	return exists, err
}

func (s *Store) ExistsByChipID(
	ctx context.Context,
	chipID string,
) (bool, error) {

	var exists bool

	err := s.db.GetContext(
		ctx,
		&exists,
		ExistsByChipID,
		chipID,
	)

	return exists, err
}

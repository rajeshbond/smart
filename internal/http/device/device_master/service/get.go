/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : get.go
 *
 * DESCRIPTION :
 * Get Device Service
 *
 ******************************************************************************/

package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// Get By ID
//------------------------------------------------------------------------------

func (s *Service) GetByID(
	ctx context.Context,
	id int64,
) (*model.Device, error) {

	return s.store.GetByID(
		ctx,
		id,
	)
}

//------------------------------------------------------------------------------
// Get By Device ID
//------------------------------------------------------------------------------

func (s *Service) GetByDeviceID(
	ctx context.Context,
	deviceID string,
) (*model.Device, error) {

	return s.store.GetByDeviceID(
		ctx,
		deviceID,
	)
}

//------------------------------------------------------------------------------
// Get By Serial Number
//------------------------------------------------------------------------------

func (s *Service) GetBySerialNumber(
	ctx context.Context,
	serialNumber string,
) (*model.Device, error) {

	return s.store.GetBySerialNumber(
		ctx,
		serialNumber,
	)
}

//------------------------------------------------------------------------------
// Get By MQTT Username
//------------------------------------------------------------------------------

func (s *Service) GetByMQTTUsername(
	ctx context.Context,
	mqttUsername string,
) (*model.Device, error) {

	return s.store.GetByMQTTUsername(
		ctx,
		mqttUsername,
	)
}

//------------------------------------------------------------------------------
// Get By Chip ID
//------------------------------------------------------------------------------

func (s *Service) GetByChipID(
	ctx context.Context,
	chipID string,
) (*model.Device, error) {

	return s.store.GetByChipID(
		ctx,
		chipID,
	)
}

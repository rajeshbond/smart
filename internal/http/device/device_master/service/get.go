/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : get.go
 *
 ******************************************************************************/

package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// Get Device By ID
//------------------------------------------------------------------------------

func (s *Service) GetByID(
	ctx context.Context,
	id int64,
) (*model.Device, error) {

	return s.Store.GetByID(ctx, id)
}

//------------------------------------------------------------------------------
// Get Device By MQTT Username
//------------------------------------------------------------------------------

func (s *Service) GetByMQTTUsername(
	ctx context.Context,
	username string,
) (*model.Device, error) {

	return s.Store.GetByMQTTUsername(ctx, username)
}

//------------------------------------------------------------------------------
// Get Device By Secret
//------------------------------------------------------------------------------

func (s *Service) GetBySecret(
	ctx context.Context,
	secret string,
) (*model.Device, error) {

	return s.Store.GetBySecret(ctx, secret)
}

//------------------------------------------------------------------------------
// Get Device By Chip ID
//------------------------------------------------------------------------------

func (s *Service) GetByChipID(
	ctx context.Context,
	chipID string,
) (*model.Device, error) {

	return s.Store.GetByChipID(ctx, chipID)
}

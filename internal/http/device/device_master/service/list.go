/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : list.go
 *
 ******************************************************************************/

package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/mapper"
)

//------------------------------------------------------------------------------
// List
//------------------------------------------------------------------------------

func (s *Service) List(
	ctx context.Context,
) ([]dto.DeviceListItem, error) {

	devices, err := s.Store.List(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ToListItems(devices), nil
}

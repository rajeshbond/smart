/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : list.go
 *
 * DESCRIPTION :
 * Device List Service
 *
 ******************************************************************************/

package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// List Devices
//------------------------------------------------------------------------------

func (s *Service) List(
	ctx context.Context,
	filter dto.DeviceFilter,
) ([]model.Device, error) {

	//----------------------------------------------------------------------
	// Default Pagination
	//----------------------------------------------------------------------

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.PageSize <= 0 {
		filter.PageSize = 20
	}

	if filter.PageSize > 100 {
		filter.PageSize = 100
	}

	//----------------------------------------------------------------------
	// Default Sorting
	//----------------------------------------------------------------------

	if filter.SortBy == "" {
		filter.SortBy = "created_at"
	}

	if filter.SortOrder == "" {
		filter.SortOrder = "DESC"
	}

	return s.store.List(
		ctx,
		filter,
	)
}

//------------------------------------------------------------------------------
// Count Devices
//------------------------------------------------------------------------------

func (s *Service) Count(
	ctx context.Context,
	filter dto.DeviceFilter,
) (int64, error) {

	return s.store.Count(
		ctx,
		filter,
	)
}

/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : update.go
 *
 ******************************************************************************/

package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/mapper"
	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// Update
//------------------------------------------------------------------------------

func (s *Service) Update(
	ctx context.Context,
	id int64,
	req dto.UpdateDeviceRequest,
	userID int64,
) (*model.Device, error) {

	//----------------------------------------------------------------------
	// Begin Transaction
	//----------------------------------------------------------------------

	tx, err := s.Store.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	//----------------------------------------------------------------------
	// DTO -> Model
	//----------------------------------------------------------------------

	device := mapper.ToModelForUpdate(
		id,
		req,
		userID,
	)

	//----------------------------------------------------------------------
	// Update
	//----------------------------------------------------------------------

	if err := s.Store.Update(ctx, tx, device); err != nil {
		return nil, err
	}

	//----------------------------------------------------------------------
	// Commit
	//----------------------------------------------------------------------

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return device, nil
}

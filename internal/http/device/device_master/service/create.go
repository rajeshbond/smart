/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : create.go
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
// Create
//------------------------------------------------------------------------------

func (s *Service) Create(
	ctx context.Context,
	req dto.CreateDeviceRequest,
	userID int64,
) (*model.Device, error) {

	//----------------------------------------------------------------------
	// Begin Transaction
	//----------------------------------------------------------------------

	tx, err := s.Store.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	//----------------------------------------------------------------------
	// DTO -> Model
	//----------------------------------------------------------------------

	device := mapper.ToModelForCreate(req, userID)

	//----------------------------------------------------------------------
	// Save
	//----------------------------------------------------------------------

	id, err := s.Store.CreateTx(
		ctx,
		tx,
		device,
	)
	if err != nil {
		return nil, err
	}

	device.ID = id

	//----------------------------------------------------------------------
	// Commit
	//----------------------------------------------------------------------

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return device, nil
}

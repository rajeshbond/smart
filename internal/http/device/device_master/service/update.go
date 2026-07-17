/******************************************************************************

 * MODULE      : Device Master
 * FILE        : update.go
 *
 * DESCRIPTION :
 * Update Device Service
 *
 ******************************************************************************/

package service

import (
	"context"
	"fmt"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// Update Device
//------------------------------------------------------------------------------

func (s *Service) Update(
	ctx context.Context,
	device *model.Device,
) error {

	//----------------------------------------------------------------------
	// Validation
	//----------------------------------------------------------------------

	if device == nil {
		return fmt.Errorf("device is nil")
	}

	if device.ID == 0 {
		return fmt.Errorf("invalid device id")
	}

	if device.DeviceID == "" {
		return fmt.Errorf("device id is required")
	}

	if device.SerialNumber == "" {
		return fmt.Errorf("serial number is required")
	}

	if device.Model == "" {
		return fmt.Errorf("model is required")
	}

	//----------------------------------------------------------------------
	// Check Device Exists
	//----------------------------------------------------------------------

	_, err := s.store.GetByID(
		ctx,
		device.ID,
	)
	if err != nil {
		return err
	}

	//----------------------------------------------------------------------
	// Update
	//----------------------------------------------------------------------

	return s.store.Update(
		ctx,
		device,
	)
}

//------------------------------------------------------------------------------
// Update Device Status
//------------------------------------------------------------------------------

func (s *Service) UpdateStatus(
	ctx context.Context,
	id int64,
	status string,
	updatedBy int64,
) error {

	if id <= 0 {
		return fmt.Errorf("invalid device id")
	}

	if status == "" {
		return fmt.Errorf("device status is required")
	}

	return s.store.UpdateStatus(
		ctx,
		id,
		status,
		updatedBy,
	)
}

//------------------------------------------------------------------------------
// Update Firmware Version
//------------------------------------------------------------------------------

func (s *Service) UpdateFirmware(
	ctx context.Context,
	id int64,
	firmwareVersion string,
	updatedBy int64,
) error {

	if id <= 0 {
		return fmt.Errorf("invalid device id")
	}

	if firmwareVersion == "" {
		return fmt.Errorf("firmware version is required")
	}

	return s.store.UpdateFirmware(
		ctx,
		id,
		firmwareVersion,
		updatedBy,
	)
}

//------------------------------------------------------------------------------
// Update Last Seen
//------------------------------------------------------------------------------

func (s *Service) UpdateLastSeen(
	ctx context.Context,
	deviceID string,
) error {

	if deviceID == "" {
		return fmt.Errorf("device id is required")
	}

	return s.store.UpdateLastSeen(
		ctx,
		deviceID,
	)
}

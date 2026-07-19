/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : create.go
 *
 * DESCRIPTION :
 * Create Device Service
 *
 ******************************************************************************/

package service

import (
	"context"
	"fmt"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// Create Device
//------------------------------------------------------------------------------

func (s *Service) Create(
	ctx context.Context,
	device *model.Device,
) (int64, error) {

	//----------------------------------------------------------------------
	// Validation
	//----------------------------------------------------------------------

	if device == nil {
		return 0, fmt.Errorf("device is nil")
	}

	if device.DeviceID == "" {
		return 0, fmt.Errorf("device id is required")
	}

	if device.SerialNumber == "" {
		return 0, fmt.Errorf("serial number is required")
	}

	if device.Model == "" {
		return 0, fmt.Errorf("model is required")
	}

	if device.MQTTUsername == "" {
		return 0, fmt.Errorf("mqtt username is required")
	}

	if device.MQTTPassword == "" {
		return 0, fmt.Errorf("mqtt password is required")
	}

	if device.DeviceSecret == "" {
		return 0, fmt.Errorf("device secret is required")
	}

	//----------------------------------------------------------------------
	// Duplicate Device ID
	//----------------------------------------------------------------------

	exists, err := s.store.ExistsByDeviceID(
		ctx,
		device.DeviceID,
	)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, fmt.Errorf("device id already exists")
	}

	//----------------------------------------------------------------------
	// Duplicate Serial Number
	//----------------------------------------------------------------------

	exists, err = s.store.ExistsBySerialNumber(
		ctx,
		device.SerialNumber,
	)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, fmt.Errorf("serial number already exists")
	}

	//----------------------------------------------------------------------
	// Duplicate MQTT Username
	//----------------------------------------------------------------------

	exists, err = s.store.ExistsByMQTTUsername(
		ctx,
		device.MQTTUsername,
	)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, fmt.Errorf("mqtt username already exists")
	}

	//----------------------------------------------------------------------
	// Duplicate Device Secret
	//----------------------------------------------------------------------

	// exists, err = s.store.ExistsByDeviceID(
	// 	ctx,
	// 	device.DeviceID,
	// )
	// if err != nil {
	// 	return 0, err
	// }

	// if exists {
	// 	return 0, fmt.Errorf("device secret already exists")
	// }

	//----------------------------------------------------------------------
	// Create Device
	//----------------------------------------------------------------------

	id, err := s.store.Create(
		ctx,
		device,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}



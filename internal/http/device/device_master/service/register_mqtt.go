/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : register_mqtt.go
 *
 ******************************************************************************/

package service

import (
	"context"
	"time"

	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// Register MQTT Username
//------------------------------------------------------------------------------

func (s *Service) RegisterMqttUsername(
	ctx context.Context,
	req dto.RegisterMQTTRequest,
	userID int64,
) (*dto.RegisterMQTTResponse, error) {

	//----------------------------------------------------------------------
	// Get Device
	//----------------------------------------------------------------------

	device, err := s.Store.GetByDeviceID(
		ctx,
		req.DeviceID,
	)
	if err != nil {
		return nil, err
	}

	//----------------------------------------------------------------------
	// Already Registered?
	//----------------------------------------------------------------------

	if device.MQTTRegistrationStatus == model.MQTTStatusRegistered {
		return nil, response.ErrMQTTAlreadyRegistered
	}

	//----------------------------------------------------------------------
	// Register User On MQTT Broker
	//----------------------------------------------------------------------

	if err = s.mqtt.RegisterUser(
		ctx,
		device.MQTTUsername,
		device.MQTTPassword,
	); err != nil {
		return nil, err
	}

	//----------------------------------------------------------------------
	// Begin Transaction
	//----------------------------------------------------------------------

	tx, err := s.Store.BeginTx(ctx)
	if err != nil {
		_ = s.mqtt.DeleteUser(ctx, device.MQTTUsername)
		return nil, err
	}

	rollback := true

	defer func() {
		if rollback {
			_ = tx.Rollback()
			_ = s.mqtt.DeleteUser(ctx, device.MQTTUsername)
		}
	}()

	//----------------------------------------------------------------------
	// Update Database
	//----------------------------------------------------------------------

	now := time.Now()

	updateReq := &dto.UpdateMQTTRegistration{

		DeviceID: device.DeviceID,

		MQTTRegistrationStatus: model.MQTTStatusRegistered,

		MQTTRegisteredAt: &now,

		MQTTRegisteredBy: &userID,

		UpdatedBy: &userID,
	}

	if err = s.Store.UpdateMQTTRegistrationTx(
		ctx,
		tx,
		updateReq,
	); err != nil {
		return nil, err
	}

	//----------------------------------------------------------------------
	// Commit Transaction
	//----------------------------------------------------------------------

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	rollback = false

	//----------------------------------------------------------------------
	// Response
	//----------------------------------------------------------------------

	return &dto.RegisterMQTTResponse{

		DeviceID: device.DeviceID,

		MQTTUsername: device.MQTTUsername,

		MQTTRegistrationStatus: model.MQTTStatusRegistered,

		MQTTRegisteredAt: &now,

		Message: "MQTT user registered successfully",
	}, nil
}

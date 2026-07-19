package service

import (
	"context"
	"time"

	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

func (s *Service) RegisterMqttUsername(
	ctx context.Context,
	req dto.RegisterMQTTRequest,
	userID int64,
) (*dto.RegisterMQTTResponse, error) {

	//----------------------------------------------------------
	// Get Device
	//----------------------------------------------------------

	device, err := s.store.GetByDeviceID(
		ctx,
		req.DeviceID,
	)

	if err != nil {
		return nil, err
	}

	//----------------------------------------------------------
	// Already Registered?
	//----------------------------------------------------------

	if device.MQTTRegistrationStatus == model.MQTTStatusRegistered {
		return nil, response.ErrMQTTAlreadyRegistered
	}

	//----------------------------------------------------------
	// Register User On Mosquitto
	//----------------------------------------------------------

	err = s.mqtt.RegisterUser(
		ctx,
		device.MQTTUsername,
		device.MQTTPassword,
	)

	if err != nil {
		return nil, err
	}

	//----------------------------------------------------------
	// Begin Transaction
	//----------------------------------------------------------

	tx, err := s.store.BeginTx(ctx)

	if err != nil {

		_ = s.mqtt.DeleteUser(
			ctx,
			device.MQTTUsername,
		)

		return nil, err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	//----------------------------------------------------------
	// Update Database
	//----------------------------------------------------------

	now := time.Now()

	updateReq := &dto.UpdateMQTTRegistration{

		DeviceID: device.DeviceID,

		MQTTRegistrationStatus: model.MQTTStatusRegistered,

		MQTTRegisteredAt: &now,

		MQTTRegisteredBy: &userID,

		UpdatedBy: &userID,
	}

	err = s.store.UpdateMQTTRegistrationTx(
		ctx,
		tx,
		updateReq,
	)

	if err != nil {

		_ = s.mqtt.DeleteUser(
			ctx,
			device.MQTTUsername,
		)

		return nil, err
	}

	//----------------------------------------------------------
	// Commit
	//----------------------------------------------------------

	err = tx.Commit()

	if err != nil {

		_ = s.mqtt.DeleteUser(
			ctx,
			device.MQTTUsername,
		)

		return nil, err
	}

	//----------------------------------------------------------
	// Response
	//----------------------------------------------------------

	resp := &dto.RegisterMQTTResponse{

		DeviceID: device.DeviceID,

		MQTTUsername: device.MQTTUsername,

		MQTTRegistrationStatus: model.MQTTStatusRegistered,

		MQTTRegisteredAt: &now,

		Message: "MQTT user registered successfully",
	}

	return resp, nil
}

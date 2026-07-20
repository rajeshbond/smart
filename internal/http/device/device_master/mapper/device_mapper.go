/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : mapper.go
 *
 ******************************************************************************/

package mapper

import (
	"github.com/rajeshbond/smart/internal/common/utils"
	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// CreateRequest -> Model
//------------------------------------------------------------------------------

func ToModelForCreate(
	req dto.CreateDeviceRequest,
	userID int64,
) *model.Device {

	return &model.Device{

		DeviceID: req.DeviceID,

		SerialNumber: req.SerialNumber,

		Model: req.Model,

		HardwareVersion: req.HardwareVersion,

		FirmwareVersion: req.FirmwareVersion,

		ManufacturedAt: req.ManufacturedAt,

		MQTTUsername: req.MQTTUsername,

		MQTTPassword: req.MQTTPassword,

		SoftAPSSID: req.SoftAPSSID,

		SoftAPPassword: req.SoftAPPassword,

		DeviceSecret: req.DeviceSecret,

		ChipID: req.ChipID,

		CommunicationType: req.CommunicationType,

		DeviceStatus: req.DeviceStatus,

		Notes: req.Notes,

		CreatedBy: &userID,

		UpdatedBy: &userID,
	}
}

//------------------------------------------------------------------------------
// UpdateRequest -> Model
//------------------------------------------------------------------------------

func ToModelForUpdate(
	id int64,
	req dto.UpdateDeviceRequest,
	userID int64,
) *model.Device {

	return &model.Device{

		ID: id,

		Model: req.Model,

		HardwareVersion: req.HardwareVersion,

		FirmwareVersion: req.FirmwareVersion,

		ManufacturedAt: req.ManufacturedAt,

		MQTTUsername: req.MQTTUsername,

		MQTTPassword: req.MQTTPassword,

		SoftAPSSID: req.SoftAPSSID,

		SoftAPPassword: req.SoftAPPassword,

		DeviceSecret: req.DeviceSecret,

		ChipID: req.ChipID,

		MACAddressWiFi: req.MacAddressWiFi,

		MACAddressEthernet: req.MacAddressEthernet,

		CommunicationType: req.CommunicationType,

		DeviceStatus: req.DeviceStatus,

		IsActive: req.IsActive,

		Notes: req.Notes,

		UpdatedBy: &userID,
	}
}

//------------------------------------------------------------------------------
// Model -> CreateResponse
//------------------------------------------------------------------------------

func ToCreateResponse(
	device *model.Device,
) dto.CreateDeviceResponse {

	return dto.CreateDeviceResponse{

		ID: device.ID,

		DeviceID: device.DeviceID,

		SerialNumber: device.SerialNumber,

		Model: device.Model,

		HardwareVersion: utils.StringValue(device.HardwareVersion),

		FirmwareVersion: utils.StringValue(device.FirmwareVersion),

		CommunicationType: device.CommunicationType,

		DeviceStatus: device.DeviceStatus,

		ChipID: utils.StringValue(device.ChipID),

		IsActive: device.IsActive,

		ManufacturedAt: utils.TimePtr(device.ManufacturedAt),

		CreatedAt: utils.TimePtr(&device.CreatedAt),

		Message: "Device created successfully",
	}
}

//------------------------------------------------------------------------------
// Model -> UpdateResponse
//------------------------------------------------------------------------------

func ToUpdateResponse(
	device *model.Device,
) dto.UpdateDeviceResponse {

	return dto.UpdateDeviceResponse{

		ID: device.ID,

		DeviceID: device.DeviceID,

		SerialNumber: device.SerialNumber,

		Model: device.Model,

		HardwareVersion: utils.StringValue(device.HardwareVersion),

		FirmwareVersion: utils.StringValue(device.FirmwareVersion),

		CommunicationType: device.CommunicationType,

		DeviceStatus: device.DeviceStatus,

		IsActive: device.IsActive,

		UpdatedAt: utils.TimePtr(&device.UpdatedAt),

		Message: "Device updated successfully",
	}
}

//------------------------------------------------------------------------------
// Model -> GetResponse
//------------------------------------------------------------------------------

func ToGetResponse(
	device *model.Device,
) dto.GetDeviceResponse {

	return dto.GetDeviceResponse{

		ID: device.ID,

		DeviceID: device.DeviceID,

		SerialNumber: device.SerialNumber,

		Model: device.Model,

		HardwareVersion: device.HardwareVersion,

		FirmwareVersion: device.FirmwareVersion,

		CommunicationType: device.CommunicationType,

		DeviceStatus: device.DeviceStatus,

		ChipID: device.ChipID,

		IsActive: device.IsActive,

		ManufacturedAt: device.ManufacturedAt,

		LastSeenAt: device.LastSeenAt,

		CreatedAt: device.CreatedAt,

		UpdatedAt: device.UpdatedAt,

		Notes: device.Notes,
	}
}

//------------------------------------------------------------------------------
// Model -> List Item
//------------------------------------------------------------------------------

func ToListItem(
	device *model.Device,
) dto.DeviceListItem {

	return dto.DeviceListItem{

		ID: device.ID,

		DeviceID: device.DeviceID,

		SerialNumber: device.SerialNumber,

		Model: device.Model,

		HardwareVersion: utils.StringValue(device.HardwareVersion),

		FirmwareVersion: utils.StringValue(device.FirmwareVersion),

		CommunicationType: device.CommunicationType,

		DeviceStatus: device.DeviceStatus,

		IsActive: device.IsActive,

		ChipID: utils.StringValue(device.ChipID),

		LastSeenAt: device.LastSeenAt,

		CreatedAt: utils.TimePtr(&device.CreatedAt),
	}
}

//------------------------------------------------------------------------------
// Model List -> DTO List
//------------------------------------------------------------------------------

func ToListItems(
	devices []*model.Device,
) []dto.DeviceListItem {

	items := make([]dto.DeviceListItem, 0, len(devices))

	for _, device := range devices {

		items = append(
			items,
			ToListItem(device),
		)
	}

	return items
}

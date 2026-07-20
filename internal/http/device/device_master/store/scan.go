/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : scan.go
 *
 ******************************************************************************/

package store

import (
	"database/sql"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

func scanDevice(scanner interface {
	Scan(dest ...any) error
}) (*model.Device, error) {

	device := &model.Device{}

	err := scanner.Scan(

		&device.ID,
		&device.DeviceID,
		&device.SerialNumber,

		&device.Model,
		&device.HardwareVersion,
		&device.FirmwareVersion,
		&device.ManufacturedAt,

		&device.MQTTUsername,
		&device.MQTTPassword,

		&device.SoftAPSSID,
		&device.SoftAPPassword,

		&device.DeviceSecret,

		&device.ChipID,

		&device.MACAddressWiFi,
		&device.MACAddressEthernet,

		&device.CommunicationType,
		&device.DeviceStatus,

		&device.LastSeenAt,

		&device.MQTTRegistrationStatus,
		&device.MQTTRegisteredAt,
		&device.MQTTRegisteredBy,

		&device.IsActive,
		&device.IsDeleted,

		&device.Notes,

		&device.CreatedAt,
		&device.CreatedBy,

		&device.UpdatedAt,
		&device.UpdatedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return device, nil
}

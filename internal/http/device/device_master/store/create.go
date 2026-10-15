/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : create.go
 *
 ******************************************************************************/

package store

import (
	"context"
	"database/sql"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

func (s *Store) CreateTx(
	ctx context.Context,
	tx *sql.Tx,
	device *model.Device,
) (int64, error) {

	var id int64

	err := tx.QueryRowContext(
		ctx,
		CreateDevice,

		device.DeviceID,
		device.SerialNumber,

		device.Model,
		device.HardwareVersion,
		device.FirmwareVersion,
		device.ManufacturedAt,

		device.MQTTUsername,
		device.MQTTPassword,

		device.SoftAPSSID,
		device.SoftAPPassword,

		device.DeviceSecret,

		device.ChipID,

		device.MACAddressWiFi,
		device.MACAddressEthernet,

		device.CommunicationType,
		device.DeviceStatus,

		device.Notes,

		device.CreatedBy,
		device.UpdatedBy,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

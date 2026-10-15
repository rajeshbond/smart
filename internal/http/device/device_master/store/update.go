/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : update.go
 *
 ******************************************************************************/

package store

import (
	"context"
	"database/sql"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

func (s *Store) Update(
	ctx context.Context,
	tx *sql.Tx,
	device *model.Device,
) error {

	result, err := tx.ExecContext(
		ctx,
		UpdateDevice,

		device.ID,

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

		device.IsActive,

		device.UpdatedBy,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

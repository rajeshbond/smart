/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : update_mqtt_registration_tx.go
 *
 ******************************************************************************/

package store

import (
	"context"
	"database/sql"

	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
)

func (s *Store) UpdateMQTTRegistrationTx(
	ctx context.Context,
	tx *sql.Tx,
	req *dto.UpdateMQTTRegistration,
) error {

	_, err := tx.ExecContext(
		ctx,
		UpdateMQTTRegistration,

		req.MQTTRegistrationStatus,
		req.MQTTRegisteredAt,
		req.MQTTRegisteredBy,
		req.UpdatedBy,
		req.DeviceID,
	)

	return err
}

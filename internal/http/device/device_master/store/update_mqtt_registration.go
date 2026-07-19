/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : update_mqtt_registration_tx.go
 *
 * DESCRIPTION :
 * Update MQTT Registration
 *
 ******************************************************************************/

package store

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
)

func (s *Store) UpdateMQTTRegistrationTx(
	ctx context.Context,
	tx *sqlx.Tx,
	req *dto.UpdateMQTTRegistration,
) error {

	_, err := tx.NamedExecContext(
		ctx,
		UpdateMQTTRegistration,
		req,
	)

	if err != nil {
		return err
	}

	return nil
}

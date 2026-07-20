/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : exists.go
 *
 ******************************************************************************/

package store

import (
	"context"
)

//------------------------------------------------------------------------------
// Exists By ID
//------------------------------------------------------------------------------

func (s *Store) ExistsByID(
	ctx context.Context,
	id int64,
) (bool, error) {

	var exists bool

	err := s.db.QueryRowContext(
		ctx,
		ExistsByID,
		id,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

//------------------------------------------------------------------------------
// Exists By Device ID
//------------------------------------------------------------------------------

func (s *Store) ExistsByDeviceID(
	ctx context.Context,
	deviceID string,
) (bool, error) {

	var exists bool

	err := s.db.QueryRowContext(
		ctx,
		ExistsByDeviceID,
		deviceID,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

//------------------------------------------------------------------------------
// Exists By Serial Number
//------------------------------------------------------------------------------

func (s *Store) ExistsBySerialNumber(
	ctx context.Context,
	serialNumber string,
) (bool, error) {

	var exists bool

	err := s.db.QueryRowContext(
		ctx,
		ExistsBySerialNumber,
		serialNumber,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

//------------------------------------------------------------------------------
// Exists By MQTT Username
//------------------------------------------------------------------------------

func (s *Store) ExistsByMQTTUsername(
	ctx context.Context,
	username string,
) (bool, error) {

	var exists bool

	err := s.db.QueryRowContext(
		ctx,
		ExistsByMQTTUsername,
		username,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

//------------------------------------------------------------------------------
// Exists By Chip ID
//------------------------------------------------------------------------------

func (s *Store) ExistsByChipID(
	ctx context.Context,
	chipID string,
) (bool, error) {

	var exists bool

	err := s.db.QueryRowContext(
		ctx,
		ExistsByChipID,
		chipID,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

//------------------------------------------------------------------------------
// Exists By Device Secret
//------------------------------------------------------------------------------

func (s *Store) ExistsByDeviceSecret(
	ctx context.Context,
	secret string,
) (bool, error) {

	var exists bool

	err := s.db.QueryRowContext(
		ctx,
		ExistsByDeviceSecret,
		secret,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

//------------------------------------------------------------------------------
// Exists By WiFi MAC Address
//------------------------------------------------------------------------------

func (s *Store) ExistsByWiFiMAC(
	ctx context.Context,
	mac string,
) (bool, error) {

	var exists bool

	err := s.db.QueryRowContext(
		ctx,
		ExistsByWiFiMAC,
		mac,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

//------------------------------------------------------------------------------
// Exists By Ethernet MAC Address
//------------------------------------------------------------------------------

func (s *Store) ExistsByEthernetMAC(
	ctx context.Context,
	mac string,
) (bool, error) {

	var exists bool

	err := s.db.QueryRowContext(
		ctx,
		ExistsByEthernetMAC,
		mac,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

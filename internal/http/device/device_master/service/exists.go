/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : exists.go
 *
 ******************************************************************************/

package service

import (
	"context"
)

//------------------------------------------------------------------------------
// Exists By ID
//------------------------------------------------------------------------------

func (s *Service) ExistsByID(
	ctx context.Context,
	id int64,
) (bool, error) {

	return s.Store.ExistsByID(ctx, id)
}

//------------------------------------------------------------------------------
// Exists By Device ID
//------------------------------------------------------------------------------

func (s *Service) ExistsByDeviceID(
	ctx context.Context,
	deviceID string,
) (bool, error) {

	return s.Store.ExistsByDeviceID(ctx, deviceID)
}

//------------------------------------------------------------------------------
// Exists By Serial Number
//------------------------------------------------------------------------------

func (s *Service) ExistsBySerialNumber(
	ctx context.Context,
	serialNumber string,
) (bool, error) {

	return s.Store.ExistsBySerialNumber(ctx, serialNumber)
}

//------------------------------------------------------------------------------
// Exists By MQTT Username
//------------------------------------------------------------------------------

func (s *Service) ExistsByMQTTUsername(
	ctx context.Context,
	username string,
) (bool, error) {

	return s.Store.ExistsByMQTTUsername(ctx, username)
}

//------------------------------------------------------------------------------
// Exists By Chip ID
//------------------------------------------------------------------------------

func (s *Service) ExistsByChipID(
	ctx context.Context,
	chipID string,
) (bool, error) {

	return s.Store.ExistsByChipID(ctx, chipID)
}

//------------------------------------------------------------------------------
// Exists By Device Secret
//------------------------------------------------------------------------------

func (s *Service) ExistsByDeviceSecret(
	ctx context.Context,
	secret string,
) (bool, error) {

	return s.Store.ExistsByDeviceSecret(ctx, secret)
}

//------------------------------------------------------------------------------
// Exists By WiFi MAC Address
//------------------------------------------------------------------------------

func (s *Service) ExistsByWiFiMAC(
	ctx context.Context,
	mac string,
) (bool, error) {

	return s.Store.ExistsByWiFiMAC(ctx, mac)
}

//------------------------------------------------------------------------------
// Exists By Ethernet MAC Address
//------------------------------------------------------------------------------

func (s *Service) ExistsByEthernetMAC(
	ctx context.Context,
	mac string,
) (bool, error) {

	return s.Store.ExistsByEthernetMAC(ctx, mac)
}

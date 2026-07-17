/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : device.go
 *
 * DESCRIPTION :
 * Device Master Model
 *
 ******************************************************************************/

package model

import "time"

type Device struct {
	ID int64 `db:"id"`

	// -------------------------------------------------------------------------
	// Device Identity
	// -------------------------------------------------------------------------

	DeviceID     string `db:"device_id"`
	SerialNumber string `db:"serial_number"`

	// -------------------------------------------------------------------------
	// Device Information
	// -------------------------------------------------------------------------

	Model           string     `db:"model"`
	HardwareVersion *string    `db:"hardware_version"`
	FirmwareVersion *string    `db:"firmware_version"`
	ManufacturedAt  *time.Time `db:"manufactured_at"`

	// -------------------------------------------------------------------------
	// Factory Provisioning
	// -------------------------------------------------------------------------

	MQTTUsername string `db:"mqtt_username"`
	MQTTPassword string `db:"mqtt_password"`

	SoftAPSSID     string `db:"softap_ssid"`
	SoftAPPassword string `db:"softap_password"`

	DeviceSecret string `db:"device_secret"`

	ChipID *string `db:"chip_id"`

	MACAddressWiFi     *string `db:"mac_address_wifi"`
	MACAddressEthernet *string `db:"mac_address_ethernet"`

	// -------------------------------------------------------------------------
	// Device Status
	// -------------------------------------------------------------------------

	CommunicationType string `db:"communication_type"`

	DeviceStatus string `db:"device_status"`

	LastSeenAt *time.Time `db:"last_seen_at"`

	IsActive bool `db:"is_active"`

	IsDeleted bool `db:"is_deleted"`

	// -------------------------------------------------------------------------
	// Additional Information
	// -------------------------------------------------------------------------

	Notes *string `db:"notes"`

	// -------------------------------------------------------------------------
	// Audit
	// -------------------------------------------------------------------------

	CreatedBy *int64 `db:"created_by"`

	UpdatedBy *int64 `db:"updated_by"`

	CreatedAt time.Time `db:"created_at"`

	UpdatedAt time.Time `db:"updated_at"`
}

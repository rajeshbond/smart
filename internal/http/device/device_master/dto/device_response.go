/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : response.go
 *
 * DESCRIPTION :
 * Device Response DTO
 *
 ******************************************************************************/

package dto

import "time"

type DeviceResponse struct {

	//-------------------------------------------------------------------------
	// Primary Key
	//-------------------------------------------------------------------------

	ID int64 `json:"id"`

	//-------------------------------------------------------------------------
	// Device Identity
	//-------------------------------------------------------------------------

	DeviceID string `json:"device_id"`

	SerialNumber string `json:"serial_number"`

	//-------------------------------------------------------------------------
	// Device Information
	//-------------------------------------------------------------------------

	Model string `json:"model"`

	HardwareVersion *string `json:"hardware_version,omitempty"`

	FirmwareVersion *string `json:"firmware_version,omitempty"`

	ManufacturedAt *time.Time `json:"manufactured_at,omitempty"`

	//-------------------------------------------------------------------------
	// Factory Provisioning
	//-------------------------------------------------------------------------

	MQTTUsername string `json:"mqtt_username"`

	SoftAPSSID string `json:"softap_ssid"`

	DeviceSecret string `json:"device_secret"`

	ChipID *string `json:"chip_id,omitempty"`

	MACAddressWiFi *string `json:"mac_address_wifi,omitempty"`

	MACAddressEthernet *string `json:"mac_address_ethernet,omitempty"`

	//-------------------------------------------------------------------------
	// Communication
	//-------------------------------------------------------------------------

	CommunicationType string `json:"communication_type"`

	DeviceStatus string `json:"device_status"`

	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`

	IsActive bool `json:"is_active"`

	//-------------------------------------------------------------------------
	// Additional Information
	//-------------------------------------------------------------------------

	Notes *string `json:"notes,omitempty"`

	//-------------------------------------------------------------------------
	// Audit
	//-------------------------------------------------------------------------

	CreatedBy *int64 `json:"created_by,omitempty"`

	UpdatedBy *int64 `json:"updated_by,omitempty"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}

/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : update.go
 *
 * DESCRIPTION :
 * Update Device DTO
 *
 ******************************************************************************/

package dto

import "time"

type UpdateDeviceRequest struct {

	//-------------------------------------------------------------------------
	// Device Information
	//-------------------------------------------------------------------------

	Model string `json:"model"`

	HardwareVersion *string `json:"hardware_version,omitempty"`

	FirmwareVersion *string `json:"firmware_version,omitempty"`

	ManufacturedAt *time.Time `json:"manufactured_at,omitempty"`

	MacAddressWiFi *string `json:"mac_address_wifi,omitempty"`

	MacAddressEthernet *string `json:"mac_address_ethernet,omitempty"`

	//-------------------------------------------------------------------------
	// Factory Provisioning
	//-------------------------------------------------------------------------

	MQTTUsername string `json:"mqtt_username"`

	MQTTPassword string `json:"mqtt_password"`

	SoftAPSSID string `json:"softap_ssid"`

	SoftAPPassword string `json:"softap_password"`

	DeviceSecret string `json:"device_secret"`

	ChipID *string `json:"chip_id,omitempty"`

	//-------------------------------------------------------------------------
	// Device
	//-------------------------------------------------------------------------

	CommunicationType string `json:"communication_type"`

	DeviceStatus string `json:"device_status"`

	IsActive bool `json:"is_active"`

	Notes *string `json:"notes,omitempty"`
}

type UpdateDeviceDTO struct {
	ID int64

	Model string

	HardwareVersion *string

	FirmwareVersion *string

	ManufacturedAt *time.Time

	MQTTUsername string

	MQTTPassword string

	SoftAPSSID string

	SoftAPPassword string

	DeviceSecret string

	ChipID *string

	MACAddressWiFi *string

	MACAddressEthernet *string

	CommunicationType string

	DeviceStatus string

	IsActive bool

	Notes *string

	UpdatedBy int64
}

type UpdateDeviceResponse struct {
	ID                int64      `json:"id"`
	DeviceID          string     `json:"device_id"`
	SerialNumber      string     `json:"serial_number"`
	Model             string     `json:"model"`
	HardwareVersion   string     `json:"hardware_version,omitempty"`
	FirmwareVersion   string     `json:"firmware_version,omitempty"`
	CommunicationType string     `json:"communication_type"`
	DeviceStatus      string     `json:"device_status"`
	IsActive          bool       `json:"is_active"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	Message           string     `json:"message"`
}

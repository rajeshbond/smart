/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : create.go
 *
 * DESCRIPTION :
 * Create Device DTO
 *
 ******************************************************************************/

package dto

import "time"

type CreateDeviceRequest struct {

	// -------------------------------------------------------------------------
	// Device Identity
	// -------------------------------------------------------------------------

	DeviceID string `json:"device_id" validate:"required"`

	SerialNumber string `json:"serial_number" validate:"required"`

	// -------------------------------------------------------------------------
	// Device Information
	// -------------------------------------------------------------------------

	Model string `json:"model" validate:"required"`

	HardwareVersion *string `json:"hardware_version,omitempty"`

	FirmwareVersion *string `json:"firmware_version,omitempty"`

	ManufacturedAt *time.Time `json:"manufactured_at,omitempty"`

	// -------------------------------------------------------------------------
	// Factory Provisioning
	// -------------------------------------------------------------------------

	MQTTUsername string `json:"mqtt_username" validate:"required"`

	MQTTPassword string `json:"mqtt_password" validate:"required"`

	SoftAPSSID string `json:"softap_ssid" validate:"required"`

	SoftAPPassword string `json:"softap_password" validate:"required"`

	DeviceSecret string `json:"device_secret" validate:"required"`

	ChipID *string `json:"chip_id,omitempty"`

	MACAddressWiFi *string `json:"mac_address_wifi,omitempty"`

	MACAddressEthernet *string `json:"mac_address_ethernet,omitempty"`

	// -------------------------------------------------------------------------
	// Device
	// -------------------------------------------------------------------------

	CommunicationType string `json:"communication_type"`

	DeviceStatus string `json:"device_status"`

	Notes *string `json:"notes,omitempty"`
}

type CreateDeviceDTO struct {
	DeviceID string

	SerialNumber string

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

	IsActive *bool

	DeviceStatus string

	Notes *string

	CreatedBy int64

	UpdatedBy int64
}

type CreateDeviceResponse struct {
	ID                int64      `json:"id"`
	DeviceID          string     `json:"device_id"`
	SerialNumber      string     `json:"serial_number"`
	Model             string     `json:"model"`
	HardwareVersion   string     `json:"hardware_version,omitempty"`
	FirmwareVersion   string     `json:"firmware_version,omitempty"`
	CommunicationType string     `json:"communication_type"`
	DeviceStatus      string     `json:"device_status"`
	ChipID            string     `json:"chip_id,omitempty"`
	IsActive          bool       `json:"is_active"`
	ManufacturedAt    *time.Time `json:"manufactured_at,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	Message           string     `json:"message"`
}

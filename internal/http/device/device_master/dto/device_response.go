package dto

import "time"

type DeviceResponse struct {
	ID int64 `json:"id"`

	// Device Identity
	DeviceID     string `json:"device_id"`
	SerialNumber string `json:"serial_number"`

	// Device Information
	Model           string  `json:"model"`
	HardwareVersion *string `json:"hardware_version,omitempty"`
	FirmwareVersion *string `json:"firmware_version,omitempty"`

	ManufacturedAt *time.Time `json:"manufactured_at,omitempty"`

	// Factory Provisioning
	MQTTUsername string `json:"mqtt_username"`

	// Never expose these
	// MQTTPassword
	// SoftAPPassword
	// DeviceSecret

	SoftAPSSID string `json:"softap_ssid"`

	ChipID *string `json:"chip_id,omitempty"`

	MACAddressWiFi     *string `json:"mac_address_wifi,omitempty"`
	MACAddressEthernet *string `json:"mac_address_ethernet,omitempty"`

	CommunicationType string `json:"communication_type"`

	DeviceStatus string `json:"device_status"`

	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`

	IsActive bool `json:"is_active"`

	IsDeleted bool `json:"is_deleted"`

	Notes *string `json:"notes,omitempty"`

	CreatedBy *int64 `json:"created_by,omitempty"`
	UpdatedBy *int64 `json:"updated_by,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

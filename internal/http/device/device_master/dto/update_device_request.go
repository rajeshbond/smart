package dto

type UpdateDeviceRequest struct {
	Model           *string `json:"model,omitempty"`
	HardwareVersion *string `json:"hardware_version,omitempty"`
	FirmwareVersion *string `json:"firmware_version,omitempty"`

	MQTTUsername *string `json:"mqtt_username,omitempty"`
	MQTTPassword *string `json:"mqtt_password,omitempty"`

	SoftAPSSID     *string `json:"softap_ssid,omitempty"`
	SoftAPPassword *string `json:"softap_password,omitempty"`

	DeviceSecret *string `json:"device_secret,omitempty"`

	ChipID *string `json:"chip_id,omitempty"`

	MACAddressWiFi     *string `json:"mac_address_wifi,omitempty"`
	MACAddressEthernet *string `json:"mac_address_ethernet,omitempty"`

	CommunicationType *string `json:"communication_type,omitempty"`

	DeviceStatus *string `json:"device_status,omitempty"`

	IsActive *bool `json:"is_active,omitempty"`

	Notes *string `json:"notes,omitempty"`
}

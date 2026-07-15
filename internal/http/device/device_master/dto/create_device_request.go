package dto

type CreateDeviceRequest struct {
	DeviceID     string `json:"device_id" validate:"required,max=50"`
	SerialNumber string `json:"serial_number" validate:"required,max=100"`

	Model           string `json:"model" validate:"required,max=100"`
	HardwareVersion string `json:"hardware_version,omitempty"`
	FirmwareVersion string `json:"firmware_version,omitempty"`

	MQTTUsername string `json:"mqtt_username" validate:"required,max=100"`
	MQTTPassword string `json:"mqtt_password" validate:"required,max=255"`

	SoftAPSSID     string `json:"softap_ssid" validate:"required,max=100"`
	SoftAPPassword string `json:"softap_password" validate:"required,max=100"`

	DeviceSecret string `json:"device_secret" validate:"required,max=255"`

	ChipID string `json:"chip_id,omitempty"`

	MACAddressWiFi     string `json:"mac_address_wifi,omitempty"`
	MACAddressEthernet string `json:"mac_address_ethernet,omitempty"`

	CommunicationType string `json:"communication_type" validate:"required,oneof=WIFI ETHERNET WIFI_ETHERNET"`

	Notes string `json:"notes,omitempty"`
}

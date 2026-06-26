package device

// RegisterRequest maps the incoming JSON payload from Next.js

type RegisterRequest struct {
	DeviceName         string `json:"device_name"`
	MQTTDeviceUsername string `json:"mqtt_device_username"`
	MQTTDevicePassword string `json:"mqtt_device_password"`
}

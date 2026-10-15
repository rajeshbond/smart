package dto

import "time"

type RegisterMQTTResponse struct {
	DeviceID               string     `json:"device_id"`
	MQTTUsername           string     `json:"mqtt_username"`
	MQTTRegistrationStatus string     `json:"mqtt_registration_status"`
	MQTTRegisteredAt       *time.Time `json:"mqtt_registered_at,omitempty"`
	Message                string     `json:"message"`
}

type UpdateMQTTRegistration struct {
	DeviceID               string     `db:"device_id"`
	MQTTRegistrationStatus string     `db:"mqtt_registration_status"`
	MQTTRegisteredAt       *time.Time `db:"mqtt_registered_at"`
	MQTTRegisteredBy       *int64     `db:"mqtt_registered_by"`
	UpdatedBy              *int64     `db:"updated_by"`
}

type RegisterMQTTRequest struct {
	DeviceID string `json:"device_id" validate:"required"`
}

package device

import (
	"context"
	"database/sql"
)

type DeviceStore struct {
	DB *sql.DB
}

func NewDeviceStore(db *sql.DB) *DeviceStore {
	return &DeviceStore{DB: db}
}

func (s *DeviceStore) CreateDevice(ctx context.Context, req RegisterRequest) error {
	query := `
	INSERT INTO device (device_name,mqtt_device_username,mqtt_device_password) VALUES ($1,$2,$3)`
	// Pass the renamed struct fields into the query parameters safely

	_, err := s.DB.ExecContext(ctx, query, req.DeviceName, req.MQTTDeviceUsername, req.MQTTDevicePassword)

	return err

}

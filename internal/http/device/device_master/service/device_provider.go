package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

type DeviceProvider interface {
	GetByDeviceID(ctx context.Context, deviceID string) (*model.Device, error)
}

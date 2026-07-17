package store

// type Store interface {
// 	Create(ctx context.Context, device *model.Device) (int64, error)
// 	Update(ctx context.Context, device *model.Device) error
// 	ExistsByDeviceID(ctx context.Context, deviceID string) (bool, error)
// 	// Delete(ctx context.Context, id int64, updatedBy int64) error
// 	Count(ctx context.Context, filter dto.DeviceFilter) (int64, error)
// }

// import (
// 	"context"

// 	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
// 	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
// )

// type Store interface {

// 	// Create

// 	Create(ctx context.Context, device *model.Device) (int64, error)

// 	// Update

// 	Update(ctx context.Context, device *model.Device) error

// 	// soft Delete

// 	Delete(ctx context.Context, id uint64) error

// 	// Get single Device

// 	GetByID(ctx context.Context, id int64) (*model.Device, error)

// 	GetDeviceByDeviceID(ctx context.Context, id string) (*model.Device, error)

// 	GetDeviceBySerialNumber(ctx context.Context, serialNumber string) (*model.Device, error)

// 	GetByChipID(ctx context.Context, chipID string) (*model.Device, error)

// 	// List

// 	List(ctx context.Context, filter dto.DeviceFilter) ([]model.Device, error)

// 	Count(ctx context.Context, filter dto.DeviceFilter) (int64, error)

// 	// Device Operation

// 	UpdateStatus(
// 		ctx context.Context,
// 		id int64,
// 		status string,
// 		updatedBy int64,
// 	) error

// 	UpdateFirmware(ctx context.Context, id int64, firmwareVerison string, updatedBy int64) error

// 	UpdateLatestSeen(ctx context.Context, deviceID string) error

// 	// Utility

// 	ExistsByDeviceID(ctx context.Context, deviceID string) (bool, error)

// 	ExistsByMQTTUserName(ctx context.Context, mqttUsername string) (bool, error)
// }

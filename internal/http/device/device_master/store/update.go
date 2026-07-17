package store

import (
	"context"
	"fmt"

	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

// Update store functions
// 1. Update Device
// 2. Update Status
// 3. Update Firmware
// 4. Update Last Seen

// Update store functions Ends

// 1. Update Device
func (s *Store) Update(ctx context.Context, device *model.Device) error {
	_, err := s.db.NamedExecContext(
		ctx,
		UpdateDevice,
		device,
	)
	if err != nil {
		return fmt.Errorf("update device: %w", err)
	}

	return nil
}

// 2. Update Statuss

func (s *Store) UpdateStatus(ctx context.Context, id int64, status string, updatedBy int64) error {

	_, err := s.db.ExecContext(
		ctx,
		UpdateDeviceStatus,
		id,
		status,
		updatedBy,
	)

	if err != nil {
		return fmt.Errorf("update device status: %w", err)
	}

	return nil
}

// 3. Update Firmware

func (s *Store) UpdateFirmware(ctx context.Context, id int64, firmwareVersion string,
	updatedBy int64) error {

	_, err := s.db.ExecContext(
		ctx,
		UpdateFirmwareVersion,
		id,
		firmwareVersion,
		updatedBy,
	)

	if err != nil {
		return fmt.Errorf("update firmware version: %w", err)
	}

	return nil
}

// 4. Update Last Seen

func (s *Store) UpdateLastSeen(ctx context.Context, deviceID string) error {

	_, err := s.db.ExecContext(
		ctx,
		UpdateLastSeen,
		deviceID,
	)

	if err != nil {
		return fmt.Errorf("update last seen: %w", err)
	}

	return nil
}

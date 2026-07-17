/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : list.go
 *
 * DESCRIPTION :
 * Device List Store
 *
 ******************************************************************************/

package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/model"
)

//------------------------------------------------------------------------------
// List Devices
//------------------------------------------------------------------------------

func (s *Store) List(
	ctx context.Context,
	filter dto.DeviceFilter,
) ([]model.Device, error) {

	//----------------------------------------------------------------------
	// Base Query
	//----------------------------------------------------------------------

	query := ListDevices

	var where []string

	var args []interface{}

	//----------------------------------------------------------------------
	// Search
	//----------------------------------------------------------------------

	if filter.Search != "" {

		search := "%" + strings.TrimSpace(filter.Search) + "%"

		where = append(where, `
(
	device_id ILIKE ?
	OR serial_number ILIKE ?
	OR model ILIKE ?
	OR mqtt_username ILIKE ?
	OR chip_id ILIKE ?
)`)

		args = append(
			args,
			search,
			search,
			search,
			search,
			search,
		)
	}

	//----------------------------------------------------------------------
	// Model
	//----------------------------------------------------------------------

	if filter.Model != "" {

		where = append(
			where,
			"model = ?",
		)

		args = append(
			args,
			filter.Model,
		)
	}

	//----------------------------------------------------------------------
	// Device Status
	//----------------------------------------------------------------------

	if filter.DeviceStatus != "" {

		where = append(
			where,
			"device_status = ?",
		)

		args = append(
			args,
			filter.DeviceStatus,
		)
	}

	//----------------------------------------------------------------------
	// Communication Type
	//----------------------------------------------------------------------

	if filter.CommunicationType != "" {

		where = append(
			where,
			"communication_type = ?",
		)

		args = append(
			args,
			filter.CommunicationType,
		)
	}

	//----------------------------------------------------------------------
	// Active
	//----------------------------------------------------------------------

	if filter.IsActive != nil {

		where = append(
			where,
			"is_active = ?",
		)

		args = append(
			args,
			*filter.IsActive,
		)
	}

	//----------------------------------------------------------------------
	// Deleted
	//----------------------------------------------------------------------

	where = append(
		where,
		"is_deleted = FALSE",
	)

	//----------------------------------------------------------------------
	// WHERE
	//----------------------------------------------------------------------

	if len(where) > 0 {

		query += "\n WHERE "

		query += strings.Join(
			where,
			"\n AND ",
		)
	}

	//----------------------------------------------------------------------
	// Sorting
	//----------------------------------------------------------------------

	sortBy := "created_at"

	switch filter.SortBy {

	case "device_id":
		sortBy = "device_id"

	case "serial_number":
		sortBy = "serial_number"

	case "model":
		sortBy = "model"

	case "device_status":
		sortBy = "device_status"

	case "communication_type":
		sortBy = "communication_type"

	case "last_seen_at":
		sortBy = "last_seen_at"

	case "updated_at":
		sortBy = "updated_at"

	case "created_at":
		sortBy = "created_at"
	}

	sortOrder := "DESC"

	if strings.EqualFold(filter.SortOrder, "ASC") {
		sortOrder = "ASC"
	}

	query += fmt.Sprintf(
		"\n ORDER BY %s %s",
		sortBy,
		sortOrder,
	)
	//----------------------------------------------------------------------
	// Pagination Defaults
	//----------------------------------------------------------------------

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.PageSize <= 0 {
		filter.PageSize = 20
	}

	if filter.PageSize > 100 {
		filter.PageSize = 100
	}

	offset := (filter.Page - 1) * filter.PageSize

	//----------------------------------------------------------------------
	// Pagination
	//----------------------------------------------------------------------

	query += `
LIMIT ?
OFFSET ?`

	args = append(
		args,
		filter.PageSize,
		offset,
	)

	//----------------------------------------------------------------------
	// PostgreSQL Rebind
	//----------------------------------------------------------------------

	query = s.db.Rebind(query)

	//----------------------------------------------------------------------
	// Execute Query
	//----------------------------------------------------------------------

	var devices []model.Device

	err := s.db.SelectContext(
		ctx,
		&devices,
		query,
		args...,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"list devices: %w",
			err,
		)
	}

	return devices, nil
}

//------------------------------------------------------------------------------
// Count Devices
//------------------------------------------------------------------------------

func (s *Store) Count(
	ctx context.Context,
	filter dto.DeviceFilter,
) (int64, error) {

	query := CountDevices

	var where []string

	var args []interface{}

	//----------------------------------------------------------------------
	// Search
	//----------------------------------------------------------------------

	if filter.Search != "" {

		search := "%" + strings.TrimSpace(filter.Search) + "%"

		where = append(where, `
(
	device_id ILIKE ?
	OR serial_number ILIKE ?
	OR model ILIKE ?
	OR mqtt_username ILIKE ?
	OR chip_id ILIKE ?
)`)

		args = append(
			args,
			search,
			search,
			search,
			search,
			search,
		)
	}

	//----------------------------------------------------------------------
	// Model
	//----------------------------------------------------------------------

	if filter.Model != "" {

		where = append(
			where,
			"model = ?",
		)

		args = append(
			args,
			filter.Model,
		)
	}

	//----------------------------------------------------------------------
	// Device Status
	//----------------------------------------------------------------------

	if filter.DeviceStatus != "" {

		where = append(
			where,
			"device_status = ?",
		)

		args = append(
			args,
			filter.DeviceStatus,
		)
	}

	//----------------------------------------------------------------------
	// Communication Type
	//----------------------------------------------------------------------

	if filter.CommunicationType != "" {

		where = append(
			where,
			"communication_type = ?",
		)

		args = append(
			args,
			filter.CommunicationType,
		)
	}

	//----------------------------------------------------------------------
	// Active
	//----------------------------------------------------------------------

	if filter.IsActive != nil {

		where = append(
			where,
			"is_active = ?",
		)

		args = append(
			args,
			*filter.IsActive,
		)
	}

	//----------------------------------------------------------------------
	// Deleted
	//----------------------------------------------------------------------

	where = append(
		where,
		"is_deleted = FALSE",
	)

	//----------------------------------------------------------------------
	// WHERE
	//----------------------------------------------------------------------

	if len(where) > 0 {

		query += "\n WHERE "

		query += strings.Join(
			where,
			"\n AND ",
		)
	}
	//----------------------------------------------------------------------
	// PostgreSQL Rebind
	//----------------------------------------------------------------------

	query = s.db.Rebind(query)

	//----------------------------------------------------------------------
	// Execute Query
	//----------------------------------------------------------------------

	var total int64

	err := s.db.GetContext(
		ctx,
		&total,
		query,
		args...,
	)

	if err != nil {
		return 0, fmt.Errorf(
			"count devices: %w",
			err,
		)
	}

	return total, nil
}

/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : list_response.go
 *
 * DESCRIPTION :
 * Device List Response DTO
 *
 ******************************************************************************/

package dto

import "time"

//------------------------------------------------------------------------------
// Device List Item
//------------------------------------------------------------------------------

type DeviceListItem struct {
	ID int64 `json:"id"`

	// Device Identity
	DeviceID     string `json:"device_id"`
	SerialNumber string `json:"serial_number"`

	// Device Information
	Model             string `json:"model"`
	HardwareVersion   string `json:"hardware_version,omitempty"`
	FirmwareVersion   string `json:"firmware_version,omitempty"`
	CommunicationType string `json:"communication_type"`

	// Status
	DeviceStatus string `json:"device_status"`
	IsActive     bool   `json:"is_active"`

	// Connectivity
	ChipID     string     `json:"chip_id,omitempty"`
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`

	// Audit
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

//------------------------------------------------------------------------------
// List Device Response
//------------------------------------------------------------------------------

type ListDeviceResponse struct {
	Items      []DeviceListItem `json:"items"`
	TotalCount int64            `json:"total_count"`
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalPages int              `json:"total_pages"`
}

package dto

type DeviceFilter struct {
	// Pagination
	Page     int `json:"page" query:"page"`
	PageSize int `json:"page_size" query:"page_size"`

	// Search
	Search string `json:"search" query:"search"`

	// Filters
	Model             string `json:"model" query:"model"`
	DeviceStatus      string `json:"device_status" query:"device_status"`
	CommunicationType string `json:"communication_type" query:"communication_type"`

	IsActive *bool `json:"is_active" query:"is_active"`

	// Sorting
	SortBy    string `json:"sort_by" query:"sort_by"`
	SortOrder string `json:"sort_order" query:"sort_order"`
}

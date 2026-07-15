package dto

type DeviceListResponse struct {
	Items []DeviceResponse `json:"items"`

	Page int `json:"page"`

	PageSize int `json:"page_size"`

	TotalRows int64 `json:"total_rows"`

	TotalPages int `json:"total_pages"`
}

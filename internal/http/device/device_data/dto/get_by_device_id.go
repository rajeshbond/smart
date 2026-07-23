package dto

import "time"

type GetProductionRequest struct {
	TenantID string `json:"tenant_id"`
	DeviceID string `json:"device_id"`
	Station  string `json:"station,omitempty"`
	Limit    int    `json:"limit"`
}

type ProductionResponse struct {
	EventID         string    `json:"event_id"`
	TenantID        string    `json:"tenant_id"`
	CustomerID      string    `json:"customer_id"`
	DeviceID        string    `json:"device_id"`
	MachineID       string    `json:"machine_id"`
	Station         string    `json:"station"`
	ProductionCount int64     `json:"production_count"`
	CycleTimeSec    float64   `json:"cycle_time_sec"`
	ProductionTime  time.Time `json:"production_time"`
	CreatedAt       time.Time `json:"created_at"`
}

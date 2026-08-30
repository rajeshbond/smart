package dto

import "time"

// ============================================================
// HEARTBEAT DTO
// ============================================================

type HeartbeatDTO struct {
	EventID   string    `json:"event_id"`
	TenantID  string    `json:"tenant_id"`
	DeviceID  string    `json:"device_id"`
	MachineID string    `json:"machine_id"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
}

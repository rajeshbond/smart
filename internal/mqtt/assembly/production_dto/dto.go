/******************************************************************************
 *
 * MODULE      : MQTT Assembly
 * FILE        : production_dto.go
 *
 ******************************************************************************/

package dto

type ProductionDTO struct {
	TenantID string `json:"tenant_id"`

	CustomerID string `json:"customer_id,omitempty"`

	DeviceID string `json:"device_id"`

	MachineID string `json:"machine_id"`

	Station string `json:"station"`

	Count uint64 `json:"count"`

	EventID string `json:"event_id"`

	CycleTimeSec float64 `json:"cycle_time_sec"`

	Timestamp string `json:"timestamp"`
}

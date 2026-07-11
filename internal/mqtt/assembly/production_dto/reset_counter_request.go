package dto

type ResetCounterRequest struct {
	TenantID string `json:"tenant_id"`

	CustomerID string `json:"customer_id"`

	DeviceID string `json:"device_id"`

	MachineID string `json:"machine_id"`

	UserID string `json:"user_id"`

	Reason string `json:"reason"`
}

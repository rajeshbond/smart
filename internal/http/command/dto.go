package command

type RestCounterRequest struct {
	TenantID   string `json:"tenant_id"`
	CustomerID string `json:"customer_id"`
	DeviceID   string `json:"device_id"`
	MachineID  string `json:"machine_id"`
	UserID     string `json:"user_id"`
	Reason     string `json:"reason"`
}

type ResetCounterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

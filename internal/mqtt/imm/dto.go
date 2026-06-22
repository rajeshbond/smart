package imm

type IMMTelemetry struct {
	TenantID    int64   `json:"tenant_id"`
	MachineID   string  `json:"machine_id"`
	Temperature float64 `json:"temperature"`
	Pressure    float64 `json:"pressure"`
	CycleTime   int     `json:"cycle_time"`
	Status      string  `json:"status"`
}

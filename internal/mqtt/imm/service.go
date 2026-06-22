package imm

type Service struct {
	Store *Store
}

func NewService(store *Store) *Service {
	return &Service{Store: store}
}

// Main telemetry Logic

func (s *Service) ProcessTelemerty(t IMMTelemetry) error {
	// Business Rules

	return s.Store.InsertTelemtry(t)
}

// heartbeat logic

func (s *Service) HandleHeartbeat(tenantID int64, machineID string) error {
	return s.Store.UpsertHeartbeat(tenantID, machineID)
}

package imm

import (
	"context"
	"log"
)

type Service struct {
	Store *Store
}

func (s *Service) ProcessTelemerty(ctx context.Context, telemetry IMMTelemetry) any {
	panic("unimplemented")
}

func NewService(store *Store) *Service {
	return &Service{Store: store}
}

// ProcessTelemetry handles main business rules and routes metrics to the persistence store
func (s *Service) ProcessTelemetry(ctx context.Context, t IMMTelemetry) error {
	// 🏭 Add any industrial business rules here (e.g., threshold validations)
	if t.CycleTime <= 0 {
		log.Printf("⚠️  [IMM Service] Received invalid cycle time (%ds) from machine", t.CycleTime)
	}

	// Forward context down to your SQL store layer
	return s.Store.InsertTelemetry(ctx, t)
}

// HandleHeartbeat tracks the connection state and watchdog pings from machines
func (s *Service) HandleHeartbeat(ctx context.Context, tenantID int64, machineID string) error {
	// Forward context down to your SQL store layer
	return s.Store.UpsertHeartbeat(ctx, tenantID, machineID)
}

// package imm

// type Service struct {
// 	Store *Store
// }

// func NewService(store *Store) *Service {
// 	return &Service{Store: store}
// }

// // Main telemetry Logic

// func (s *Service) ProcessTelemerty(t IMMTelemetry) error {
// 	// Business Rules

// 	return s.Store.InsertTelemtry(t)
// }

// // heartbeat logic

// func (s *Service) HandleHeartbeat(tenantID int64, machineID string) error {
// 	return s.Store.UpsertHeartbeat(tenantID, machineID)
// }

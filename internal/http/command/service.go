package command

import (
	"context"
	"time"

	"github.com/rajeshbond/smart/internal/mqtt"
	"github.com/rajeshbond/smart/internal/mqtt/assembly"
	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
)

type CommandService struct {
	publisher *mqtt.Publisher
}

func NewCommandService(publisher *mqtt.Publisher) *CommandService {
	return &CommandService{publisher: publisher}
}

func (s *CommandService) ResetCounter(
	ctx context.Context,
	req dto.ResetCounterRequest,
) error {
	command := dto.CommandDTO{
		Command:    "RESET_COUNTER",
		TenantID:   req.TenantID,
		CustomerID: req.CustomerID,
		DeviceID:   req.DeviceID,
		MachineID:  req.MachineID,
		UserID:     req.UserID,
		Reason:     req.Reason,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	topic := assembly.TopicCommand + req.DeviceID

	return s.publisher.Publish(topic, command)
}

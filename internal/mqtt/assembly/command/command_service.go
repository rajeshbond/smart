package command

// type CommandSerivce struct {
// 	publisher *mqtt.Publisher
// }

// func NewCommandService(pub *mqtt.Publisher) *CommandSerivce {
// 	return &CommandSerivce{publisher: pub}
// }

// func (s *CommandSerivce) ResetCounter(
// 	tenantID string,
// 	customerID string,
// 	deviceID string,
// 	machineID string,
// 	userID string,
// 	reason string,
// ) error {
// 	cmd := dto.CommandDTO{
// 		Command:    "RESET_COUNTER",
// 		TenantID:   tenantID,
// 		CustomerID: customerID,
// 		DeviceID:   deviceID,
// 		MachineID:  machineID,
// 		UserID:     userID,
// 		Reason:     reason,
// 		Timestamp:  time.Now().Format(time.RFC3339),
// 	}

// 	topic := fmt.Sprintf("%s %s", assembly.TopicCommand, deviceID)

// 	return s.publisher.Publish(topic, cmd)
// }

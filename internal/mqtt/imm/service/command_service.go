package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	dto "github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

// ============================================================
// MQTT COMMAND TOPIC
// ============================================================

const TopicCommand = "factory/imm/command/"

// ============================================================
// COMMAND SERVICE
// ============================================================

type CommandService struct {
	mqttClient mqtt.Client
}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewCommandService(
	mqttClient mqtt.Client,
) *CommandService {

	return &CommandService{
		mqttClient: mqttClient,
	}
}

// ============================================================
// PUBLISH COMMAND
// ============================================================

func (s *CommandService) PublishCommand(
	deviceID string,
	req dto.IMMCommandRequest,
) (string, error) {

	// --------------------------------------------------------
	// VALIDATE DEVICE ID
	// --------------------------------------------------------

	deviceID = strings.TrimSpace(deviceID)

	if deviceID == "" {
		return "", fmt.Errorf(
			"device_id is required",
		)
	}

	// --------------------------------------------------------
	// VALIDATE COMMAND
	// --------------------------------------------------------

	req.Command = strings.TrimSpace(req.Command)

	if req.Command == "" {
		return "", fmt.Errorf(
			"command is required",
		)
	}

	// --------------------------------------------------------
	// CHECK MQTT CLIENT
	// --------------------------------------------------------

	if s.mqttClient == nil {
		return "", fmt.Errorf(
			"MQTT client is nil",
		)
	}

	// --------------------------------------------------------
	// CHECK MQTT CONNECTION
	// --------------------------------------------------------

	if !s.mqttClient.IsConnected() {
		return "", fmt.Errorf(
			"MQTT broker is not connected",
		)
	}

	// --------------------------------------------------------
	// BUILD DEVICE-SPECIFIC TOPIC
	// --------------------------------------------------------

	topic := TopicCommand + deviceID

	// Example:
	//
	// deviceID = 04@xoom
	//
	// topic =
	// factory/imm/command/04@xoom

	// --------------------------------------------------------
	// CREATE JSON PAYLOAD
	// --------------------------------------------------------

	payload, err := json.Marshal(req)

	if err != nil {
		return "", fmt.Errorf(
			"failed to create command JSON: %w",
			err,
		)
	}

	// --------------------------------------------------------
	// MQTT PUBLISH
	// --------------------------------------------------------
	//
	// QoS    = 1
	// Retain = false
	//
	// IMPORTANT:
	// Commands should NOT be retained.
	//

	token := s.mqttClient.Publish(
		topic,
		byte(1),
		false,
		payload,
	)

	// --------------------------------------------------------
	// WAIT FOR PUBLISH
	// --------------------------------------------------------

	if !token.Wait() {
		return "", fmt.Errorf(
			"MQTT publish timeout",
		)
	}

	// --------------------------------------------------------
	// CHECK PUBLISH ERROR
	// --------------------------------------------------------

	if err := token.Error(); err != nil {
		return "", fmt.Errorf(
			"MQTT publish failed: %w",
			err,
		)
	}

	// --------------------------------------------------------
	// LOG
	// --------------------------------------------------------

	log.Printf(
		"📤 IMM COMMAND PUBLISHED | Device=%s | Topic=%s | Payload=%s",
		deviceID,
		topic,
		string(payload),
	)

	return topic, nil
}

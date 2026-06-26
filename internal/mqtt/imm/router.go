package imm

import (
	paho "github.com/eclipse/paho.mqtt.golang"
)

func (m *Module) MQTTRoute(client paho.Client, subscribe func(paho.Client, string, paho.MessageHandler)) {
	// Track raw

	subscribe(client, "imm/+/telemetry", m.Handler.TelemetryHandler())
}

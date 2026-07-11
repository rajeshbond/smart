package imm

import (
	paho "github.com/eclipse/paho.mqtt.golang"
)

func (m *Module) RegisterMQTTIMMRoute(
	client paho.Client,
	subscribe func(paho.Client, string, paho.MessageHandler),
) {
	subscribe(client, "imm/+/telemetry", m.Handler.TelemetryHandler())
}

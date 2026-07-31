package assembly

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
)

func (m *Module) RegisterMQTTAssemblyRoute(client paho.Client,
	subscribe func(paho.Client, string, paho.MessageHandler)) {
	log.Println("------------------------------------------------")
	log.Println("Registering Assembly MQTT Routes...")
	log.Println("------------------------------------------------")

	// Production

	subscribe(client, TopicProduction, m.Handler.ProductionHandler())
	subscribe(client, TopicHeartbeat, m.Handler.HeartBeatHandler())

	// Heartbeat
	// subscribe(client,TopicHeartbeat,m.Handler.ProductionHandler())

	log.Println("Assembly MQTT Routes Registered")
}

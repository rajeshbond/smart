package assembly

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// ============================================================
// REGISTER MQTT ROUTES
// ============================================================

func (m *Module) RegisterMQTTAssemblyRoute(
	client paho.Client,
	subscribe func(
		paho.Client,
		string,
		paho.MessageHandler,
	),
) {

	log.Println("------------------------------------------------")
	log.Println("Registering Assembly MQTT Routes...")
	log.Println("------------------------------------------------")

	// ========================================================
	// PRODUCTION
	// ========================================================

	subscribe(
		client,
		TopicProduction,
		m.Handlers.Production.ProductionHandler(),
	)

	log.Printf(
		"✅ Production Route Registered | Topic=%s",
		TopicProduction,
	)

	// ========================================================
	// HEARTBEAT
	// ========================================================

	subscribe(
		client,
		TopicHeartbeat,
		m.Handlers.Heartbeat.HeartBeatHandler(),
	)

	log.Printf(
		"✅ Heartbeat Route Registered | Topic=%s",
		TopicHeartbeat,
	)

	log.Println("------------------------------------------------")
	log.Println("Assembly MQTT Routes Registered")
	log.Println("------------------------------------------------")
}

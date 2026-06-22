package mqtt

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/rajeshbond/smart/config"
	"github.com/rajeshbond/smart/database"
)

func Start(
	db *database.DB,
	cfg *config.Config,
) paho.Client {
	client := NewClient(cfg)
	RegisterRoutes(client, db)

	return client
}

// package mqtt

// import (
// 	"time"

// 	"github.com/rajeshbond/smart/config"
// 	"github.com/rajeshbond/smart/database"
// 	imm "github.com/rajeshbond/smart/internal/mqtt/imm"

// 	paho "github.com/eclipse/paho.mqtt.golang"
// )

// func Start(db *database.DB, cfg *config.Config) paho.Client {

// 	client := NewClient(cfg)

// 	// telemetry
// 	immService := imm.NewMQTTService(db)

// 	client.Subscribe(imm.TopicTelemetry, 1, imm.MessageHandler(immService))

// 	// heartbeat watchdog handler
// 	heartbeatService := imm.NewWatchdogService(db)

// 	client.Subscribe("tenant/+/imm/+/heartbeat", 1, heartbeatService.Handler())

// 	// OFFLINE checker
// 	StartWatchdog(db, 30*time.Second)

// 	return client
// }

// func Start(db *database.DB, cfg *config.Config) paho.Client {

// 	client := NewClient(cfg)

// 	// telemetry
// 	immService := imm.NewMQTTService(db)

// 	client.Subscribe(imm.TopicTelemetry, 1, imm.MessageHandler(immService))

// 	// heartbeat watchdog handler
// 	heartbeatService := imm.NewWatchdogService(immService)

// 	client.Subscribe("tenant/+/imm/+/heartbeat", 1, heartbeatService.Handler())

// 	// OFFLINE checker
// 	StartWatchdog(db, 30*time.Second)

// 	return client
// }

// package mqtt

// import (
// 	"log"

// 	"github.com/rajeshbond/smart/config"
// 	"github.com/rajeshbond/smart/database"
// 	imm "github.com/rajeshbond/smart/internal/mqtt/imm"

// 	paho "github.com/eclipse/paho.mqtt.golang"
// )

// func Start(db *database.DB, cfg *config.Config) paho.Client {

// 	client := NewClient(cfg)

// 	immService := imm.NewMQTTService(db)

// 	token := client.Subscribe(
// 		imm.TopicTelemetry,
// 		1,
// 		imm.MessageHandler(immService),
// 	)

// 	token.Wait()

// 	if token.Error() != nil {
// 		log.Fatal(token.Error())
// 	}

// 	log.Println("📡 MQTT Started & Subscribed")

// 	return client
// }

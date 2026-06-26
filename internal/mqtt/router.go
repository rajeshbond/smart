package mqtt

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/rajeshbond/smart/database"
	"github.com/rajeshbond/smart/internal/mqtt/imm"
)

// 1. Passively mount incoming data streams
func RegisterRoutes(client paho.Client, db *database.DB) {

	// --------------- 1. IMM Module -------------------
	imm.NewModule(db).MQTTRoute(client, subscribe)

}

// 2. Actively boot up your background automated sweeping jobs
func RegisterWorkers(db *database.DB) {
	log.Println("🐕 Initializing background factory hardware watchdogs...")

	// Boot up IMM Watchdog (30 Second Timeout)
	// immStore := imm.NewStore(db.PGX)
	// imm.StartWatchdog(immStore, 30*time.Second)

	// Boot up Assembly Watchdog if needed (1 Minute Timeout)
	// assemblyStore := assembly.NewStore(db)
	// assembly.StartWatchdog(assemblyStore, 60*time.Second)

	log.Println("✅ All background factory watchdogs are active")
}
func subscribe(client paho.Client, topic string, handler paho.MessageHandler) {
	token := client.Subscribe(topic, 1, handler)
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("❌ Subscription failed for [%s]: %v", topic, token.Error())
	}
}

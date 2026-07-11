package mqtt

import (
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"

	"github.com/rajeshbond/smart/config"
	"github.com/rajeshbond/smart/database"
)

func NewClient(
	cfg *config.Config,
	db *database.DB,
) paho.Client {

	opts := paho.NewClientOptions()

	opts.AddBroker(cfg.MQTTBROKER)
	opts.SetClientID(cfg.MQTTCLIENTID)
	opts.SetUsername(cfg.MQTTUSERNAME)
	opts.SetPassword(cfg.MQTTPASSWORD)

	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(15 * time.Second)
	opts.SetKeepAlive(30 * time.Second)

	// Optional but recommended
	opts.SetCleanSession(false)

	// Called after FIRST connect and EVERY reconnect
	opts.OnConnect = func(client paho.Client) {

		log.Println("✅ MQTT Connected")

		RegisterRoutes(client, db)

		log.Println("✅ MQTT Routes Registered")
	}

	opts.OnConnectionLost = func(client paho.Client, err error) {

		log.Printf("❌ MQTT Connection Lost : %v", err)

	}

	return paho.NewClient(opts)

}

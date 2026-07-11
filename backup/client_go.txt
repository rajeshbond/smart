package mqtt

import (
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/rajeshbond/smart/config"
)

func NewClient(cfg *config.Config) paho.Client {

	opts := paho.NewClientOptions()
	opts.AddBroker(cfg.MQTTBROKER)
	opts.SetClientID(cfg.MQTTCLIENTID)
	opts.SetUsername(cfg.MQTTUSERNAME)
	opts.SetPassword(cfg.MQTTPASSWORD)

	// Clean recovery behaviours

	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(15 * time.Second)
	opts.SetKeepAlive(30 * time.Second)

	opts.OnConnect = func(c paho.Client) {
		log.Println("🔄 MQTT Connection established with broker; registering routes...")
	}

	opts.OnConnectionLost = func(c paho.Client, err error) {
		log.Printf("❌ MQTT Connection lost: %v", err)
	}

	return paho.NewClient(opts)

}

package mqtt

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/rajeshbond/smart/config"
)

func NewClient(cfg *config.Config) paho.Client {

	opts := paho.NewClientOptions()

	opts.AddBroker(cfg.MQTTBROKER)
	opts.SetClientID(cfg.MQTTCLIENTID)

	opts.SetUsername(cfg.MQTTUSERNAME)
	opts.SetPassword(cfg.MQTTPASSWORD)

	opts.AutoReconnect = true
	opts.ConnectRetry = true

	client := paho.NewClient(opts)

	token := client.Connect()
	token.Wait()

	if token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Println("✅ MQTT Connected")

	return client
}

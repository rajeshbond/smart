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

// func NewClient(cfg *config.Config) paho.Client {
// 	fmt.Println("------> User Name:", cfg.MQTTUSERNAME)
// 	fmt.Println("------> Password :", cfg.MQTTPASSWORD)
// 	opts := paho.NewClientOptions()
// 	opts.AddBroker(cfg.MQTTBROKER)
// 	opts.SetClientID(cfg.MQTTCLIENTID)
// 	opts.SetUsername(cfg.MQTTUSERNAME)
// 	opts.SetPassword(cfg.MQTTPASSWORD)

// 	opts.AutoReconnect = true
// 	opts.ConnectRetry = true

// 	return paho.NewClient(opts)

// }

// func NewClient(cfg *config.Config) paho.Client {

// 	opts := paho.NewClientOptions()
// 	fmt.Println("------> User Name:", cfg.MQTTUSERNAME)
// 	fmt.Println("------> User Name:", cfg.MQTTPASSWORD)
// 	opts.AddBroker(cfg.MQTTBROKER)
// 	opts.SetClientID(cfg.MQTTCLIENTID)

// 	opts.SetUsername(cfg.MQTTUSERNAME)
// 	opts.SetPassword(cfg.MQTTPASSWORD)

// 	opts.AutoReconnect = true
// 	opts.ConnectRetry = true

// 	client := paho.NewClient(opts)

// 	token := client.Connect()
// 	token.Wait()

// 	if token.Error() != nil {
// 		log.Fatal(token.Error())
// 	}

// 	log.Println("✅ MQTT Connected")

// 	return client
// }

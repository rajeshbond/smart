package mqtt

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
)

type SubscribeFunc func(
	client paho.Client,
	topic string,
	handler paho.MessageHandler,
)

func subscribe(
	client paho.Client,
	topic string,
	handler paho.MessageHandler,
) {

	token := client.Subscribe(
		topic,
		1,
		handler,
	)

	token.Wait()

	if err := token.Error(); err != nil {
		log.Printf(
			"❌ Failed subscribing [%s]: %v",
			topic,
			err,
		)
		return
	}

	log.Printf("✅ MQTT Subscribed : %s", topic)
}

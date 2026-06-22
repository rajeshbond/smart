package mqtt

import (
	"encoding/json"
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
)

type Publisher struct {
	Client paho.Client
}

func NewPublisher(client paho.Client) *Publisher {
	return &Publisher{
		Client: client,
	}
}

// Publish any struct as JSON
func (p *Publisher) Publish(topic string, payload any) error {

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	token := p.Client.Publish(
		topic,
		1,     // QoS
		false, // retain
		data,
	)

	token.Wait()

	if token.Error() != nil {
		return token.Error()
	}

	log.Println("📤 Published to:", topic)

	return nil
}

package mqtt

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/rajeshbond/smart/database"
	"github.com/rajeshbond/smart/internal/mqtt/imm"
)

func RegisterRoutes(client paho.Client, db *database.DB) {

	// IMM

	immStore := imm.NewStore(db)
	immService := imm.NewService(immStore)
	immHandler := imm.NewHandler(immService)

	subscribe(
		client,
		imm.TopicTelemrty,
		immHandler.TelemetryHandler(),
	)

}

func subscribe(client paho.Client, topic string, handler paho.MessageHandler) {
	token := client.Subscribe(
		topic, 1, handler,
	)
	token.Wait()
	if token.Error() != nil {
		log.Fatal((token.Error()))
	}
}

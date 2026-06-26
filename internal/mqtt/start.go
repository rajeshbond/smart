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

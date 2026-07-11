package application

import (
	"log"
	"net/http"

	paho "github.com/eclipse/paho.mqtt.golang"

	"github.com/rajeshbond/smart/config"
	"github.com/rajeshbond/smart/database"
	"github.com/rajeshbond/smart/internal/mqtt"
)

type App struct {
	DB         *database.DB
	Config     *config.Config
	MQTTClient paho.Client
}

func NewApp() *App {

	cfg := config.Load()

	log.Println("1. Loading Database...")

	db := database.NewDB(cfg)

	log.Println("2. Initializing MQTT Client...")

	mqttClient := mqtt.NewClient(
		cfg,
		db,
	)

	log.Println("3. NewApp setup complete!")

	return &App{
		DB:         db,
		Config:     cfg,
		MQTTClient: mqttClient,
	}
}

func (a *App) Start() error {

	r := NewRouter(a)

	go func() {

		log.Println("🔄 Connecting MQTT...")

		token := a.MQTTClient.Connect()

		token.Wait()

		if err := token.Error(); err != nil {

			log.Printf("⚠️ MQTT Connection Failed : %v", err)

			return
		}

		log.Println("✅ MQTT Initial Connection Successful")

		// Start background workers only once
		mqtt.RegisterWorkers(a.DB)

	}()

	log.Println("🚀 HTTP Server Running :", a.Config.APPPORT)

	return http.ListenAndServe(
		":"+a.Config.APPPORT,
		r,
	)

}

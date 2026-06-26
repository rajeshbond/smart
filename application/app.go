package application

import (
	"log"
	"net/http"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/rajeshbond/smart/config"
	"github.com/rajeshbond/smart/database"
	"github.com/rajeshbond/smart/internal/mqtt" // Contains both RegisterRoutes and RegisterWorkers
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

	log.Println("2. Initializing MQTT client options...")
	mqttClient := mqtt.NewClient(cfg)
	log.Println("3. NewApp setup complete!")

	return &App{
		DB:         db,
		Config:     cfg,
		MQTTClient: mqttClient,
	}
}

func (a *App) Start() error {
	// 1. Initialize your HTTP Application Web Routes (Chi Router Layout)
	r := NewRouter(a)

	// 2. Launch the background MQTT initialization network routine
	go func() {
		log.Println("🔄 Initializing background MQTT handshake...")
		token := a.MQTTClient.Connect()

		if token.Wait() && token.Error() != nil {
			log.Printf("⚠️ MQTT Connection deferred: %v (Auto-retrying in background)", token.Error())
			return
		}

		log.Println("✅ MQTT Connected Successfully")

		// 🛰️ 1. PASSIVE NETWORK INBOUND ROUTING
		// Mounts IMM, Assembly, and PressShop message routes cleanly
		mqtt.RegisterRoutes(a.MQTTClient, a.DB)

		// 🐕 2. ACTIVE BACKGROUND DATABASE AUTOMATION WORKERS
		// Boots up all single-query background cron sweeps (watchdogs) for your modules
		mqtt.RegisterWorkers(a.DB)
	}()

	// 🚀 Launch the blocking HTTP Web Server on the main thread
	log.Println("🚀 HTTP Server running on :", a.Config.APPPORT)
	return http.ListenAndServe(":"+a.Config.APPPORT, r)
}

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

	// Pre-configures MQTT client, username, and options but does NOT connect yet
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
	// Build your HTTP application routes

	r := NewRouter(a)

	go func() {
		log.Println("🔄 Initializing background MQTT handshake...")
		token := a.MQTTClient.Connect()

		// 🛠️ FIX: Removed the malformed duplicate if-statement check that broke compilation
		if token.Wait() && token.Error() != nil {
			log.Printf("⚠️ MQTT Connection deferred: %v (Auto-retrying in background)", token.Error())
			return
		}

		log.Println("✅ MQTT Connected Successfully")
	}()

	// go func() {
	// 	log.Println("🔄 Initializing background MQTT handshake...")
	// 	token := a.MQTTClient.Connect()

	// 	// ⏱️ Wait for a maximum of 5 seconds for the broker to respond
	// 	// Make sure to import "time" at the top of your file
	// 	completed := token.WaitTimeout(5 * time.Second)

	// 	if !completed {
	// 		log.Println("❌ MQTT Connection Timeout: The broker did not respond within 5 seconds. Is Mosquitto running?")
	// 		return
	// 	}

	// 	if token.Error() != nil {
	// 		log.Printf("❌ MQTT Connection Rejected: %v", token.Error())
	// 		return
	// 	}

	// 	log.Println("✅ MQTT Connected Successfully")
	// }()

	// 🚀 Path B: Launch the HTTP Web Server on the main thread
	log.Println("🚀 HTTP Server running on :", a.Config.APPPORT)

	// This action blocks the main thread so your application stays alive
	return http.ListenAndServe(":"+a.Config.APPPORT, r)

}

// import (
// 	"log"
// 	"net/http"

// 	paho "github.com/eclipse/paho.mqtt.golang"
// 	"github.com/rajeshbond/smart/config"
// 	"github.com/rajeshbond/smart/database"
// 	"github.com/rajeshbond/smart/internal/mqtt"
// )

// type App struct {
// 	DB         *database.DB
// 	Config     *config.Config
// 	MQTTClient paho.Client
// }

// func (a *App) ConnectMQTT() {
// 	panic("unimplemented")
// }

// func (a *App) StartServer() any {
// 	panic("unimplemented")
// }

// func NewApp() *App {
// 	cfg := config.Load()
// 	db := database.NewDB(cfg)
// 	// Start MQTT
// 	mqttClient := mqtt.Start(db, cfg)

// 	return &App{
// 		DB:         db,
// 		Config:     cfg,
// 		MQTTClient: mqttClient,
// 	}

// }

// func (a *App) Start() error {
// 	r := NewRouter(a)
// 	log.Println("🚀 HTTP Server running on :", a.Config.APPPORT)

// 	return http.ListenAndServe(":"+a.Config.APPPORT, r)

// }

// package application

// import (
// 	"log"
// 	"net/http"

// 	"github.com/rajeshbond/smart/config"
// 	"github.com/rajeshbond/smart/database"
// )

// type App struct {
// 	DB     *database.DB
// 	Config *config.Config
// }

// func NewApp() *App {

// 	cfg := config.Load()
// 	db := database.NewDB(cfg)

// 	return &App{
// 		DB:     db,
// 		Config: cfg,
//databadatabase.
// }

// func (a *App) Start() error {

// 	r := NewRouter(a)

// 	log.Println("🚀 HTTP Server running on:", a.Config.APPPORT)

// 	return http.ListenAndServe(":"+a.Config.APPPORT, r)
// }

// package application

// import (
// 	"log"
// 	"net/http"

// 	"github.com/rajeshbond/smart/config"
// 	"github.com/rajeshbond/smart/database"
// )

// type App struct {
// 	DB     *database.DB
// 	Config *config.Config
// }

// func NewApp() *App {
// 	cfg := config.Load()
// 	db := database.NewDB(cfg)

// 	return &App{
// 		DB:     db,
// 		Config: cfg,
// 	}
// }

// func (a *App) Start() error {
// 	defer func() {
// 		if a.DB != nil {
// 			a.DB.Close()
// 		}
// 	}()

// 	r := NewRouter(a)

// 	log.Println("🚀 Server running on:", a.Config.APPPORT)

// 	return http.ListenAndServe(":"+a.Config.APPPORT, r)

// }

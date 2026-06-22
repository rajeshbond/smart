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
	db := database.NewDB(cfg)
	// Start MQTT
	mqttClient := mqtt.Start(db, cfg)

	return &App{
		DB:         db,
		Config:     cfg,
		MQTTClient: mqttClient,
	}

}

func (a *App) Start() error {
	r := NewRouter(a)
	log.Println("🚀 HTTP Server running on :", a.Config.APPPORT)

	return http.ListenAndServe(":"+a.Config.APPPORT, r)

}

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
//database
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

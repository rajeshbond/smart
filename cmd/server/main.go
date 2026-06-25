package main

import (
	"log"

	"github.com/rajeshbond/smart/application"
	"github.com/rajeshbond/smart/internal/common/utils"
)

func main() {
	utils.InitValidator()

	app := application.NewApp()

	// Handler grecefull connection diconnect on shutdowm

	defer func() {
		if app.MQTTClient != nil && app.MQTTClient.IsConnected() {
			app.MQTTClient.Disconnect(250)
			log.Println("🔌 MQTT disconnected cleanly")
		}
		if app.DB != nil {
			app.DB.Close()
			log.Println("🔌 Database pool closed cleanly")
		}
	}()

	// Starts HTTP immediately and hands MQTT over to a background thread
	if err := app.Start(); err != nil {
		log.Fatal("❌ Critical Server Failure: ", err)
	}
}

// func main() {
// 	utils.InitValidator()

// 	app := application.NewApp()

// 	defer func() {
// 		if app.MQTTClient != nil && app.MQTTClient.IsConnected() {
// 			app.MQTTClient.Disconnect(250)

// 			if app.DB != nil {
// 				app.DB.Close()
// 			}
// 		}
// 	}()

// 	if err := app.Start(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// package main

// import (
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/rajeshbond/smart/application"
// 	"github.com/rajeshbond/smart/internal/common/utils"
// 	"github.com/rajeshbond/smart/internal/mqtt"
// )

// func main() {

// 	utils.InitValidator()

// 	app := application.NewApp()

// 	// Start MQTT first (backgrutils.V
// 	mqttClient := mqtt.Start(app.DB, app.Config)

// 	// Run HTTP server in background
// 	serverErr := make(chan error, 1)
// 	go func() {
// 		serverErr <- app.Start()
// 	}()

// 	// Wait for shutdown signal
// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

// 	select {
// 	case err := <-serverErr:
// 		log.Fatal("HTTP server error:", err)

// 	case sig := <-sigChan:
// 		log.Println("Shutdown signal received:", sig)
// 	}

// 	// CLEAN SHUTDOWN
// 	log.Println("Stopping services...")

// 	if mqttClient != nil && mqttClient.IsConnected() {
// 		mqttClient.Disconnect(250)
// 		log.Println("MQTT disconnected")
// 	}

// 	app.DB.Close()

// 	log.Println("Server stopped cleanly")
// }

// package main

// import (
// 	"log"

// 	"github.com/rajeshbond/smart/application"
// 	"github.com/rajeshbond/smart/internal/common/utils"
// )

// func main() {
// 	utils.InitValidator()
// 	app := application.NewApp()
// 	defer app.DB.Close()
// 	if err := app.Start(); err != nil {
// 		log.Fatal(err)
// 	}
// }

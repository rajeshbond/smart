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

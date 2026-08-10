package handler

import (
	"context"
	"encoding/json"
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
)

func (h *ProductionHandler) HeartBeatHandler() paho.MessageHandler {
	return func(client paho.Client, msg paho.Message) {

		// Deep-copy payload safely before passing to background goroutine
		payload := make([]byte, len(msg.Payload()))
		copy(payload, msg.Payload())

		go func(data []byte) {

			// Print complete raw MQTT payload
			// log.Printf("========== COMPLETE MQTT PAYLOAD ==========")
			// log.Printf("%s", string(data))
			// log.Printf("===========================================")

			var req dto.ProductionDTO

			if err := json.Unmarshal(data, &req); err != nil {
				log.Printf("Production JSON Error: %v", err)
				log.Printf("Raw Payload: %q", string(data))
				return
			}

			_, cancel := context.WithTimeout(
				context.Background(),
				5*time.Second,
			)
			defer cancel()

			log.Println("=============== Heart Beat ===============")
			log.Printf("Heart Beat DTO: %+v", req)
			log.Println("===========================================")

		}(payload)
	}
}

// Back up code

// func (h *ProductionHandler) HeartBeatHandler() paho.MessageHandler {
// 	return func(client paho.Client, msg paho.Message) {
// 		// Deep-copy payload safely before passing to background goroutine
// 		payload := make([]byte, len(msg.Payload()))
// 		copy(payload, msg.Payload())

// 		go func(data []byte) {
// 			var req dto.ProductionDTO

// 			if err := json.Unmarshal(data, &req); err != nil {
// 				log.Printf("Production JSON Error : %v", err)
// 				// FIXED: Use 'data' instead of 'msg.Payload()' to prevent data race
// 				log.Printf("Raw Payload: %q", string(data))
// 				return
// 			}

// 			_, cancel := context.WithTimeout(
// 				context.Background(),
// 				5*time.Second,
// 			)
// 			defer cancel()

// 			log.Println("===============Heart Beat===================")
// 			log.Println("===Heart Beat====>", req)
// 			log.Println("============================================")

// 		}(payload)
// 	}
// }

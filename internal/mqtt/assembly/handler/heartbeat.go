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
			var req dto.ProductionDTO

			if err := json.Unmarshal(data, &req); err != nil {
				log.Printf("Production JSON Error : %v", err)
				// FIXED: Use 'data' instead of 'msg.Payload()' to prevent data race
				log.Printf("Raw Payload: %q", string(data))
				return
			}

			_, cancel := context.WithTimeout(
				context.Background(),
				5*time.Second,
			)
			defer cancel()

			log.Println("==================================")
			log.Println("===Heart Beat====>", req)
			log.Println("==================================")
			// ctx, cancel := context.WithTimeout(
			// 	context.Background(),
			// 	5*time.Second,
			// )
			// defer cancel()

			// if err := h.service.Save(ctx, &req); err != nil {
			// 	log.Printf("Save Error : %v", err)
			// 	return
			// }

			// log.Printf(
			// 	"Production Saved | EventID=%s | Device=%s | Station=%s | Count=%d | Cycle Time=%.2f sec",
			// 	req.EventID,
			// 	req.DeviceID,
			// 	req.Station,
			// 	req.Count,
			// 	req.CycleTimeSec,
			// )
		}(payload)
	}
}

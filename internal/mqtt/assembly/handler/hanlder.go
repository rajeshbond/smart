package handler

import (
	"context"
	"encoding/json"
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
	"github.com/rajeshbond/smart/internal/mqtt/assembly/service"
)

type ProductionHandler struct {
	service service.ProductionSerive
}

func NewProductionHandler(service *service.ProductionSerive) *ProductionHandler {
	return &ProductionHandler{service: *service}
}

func (h *ProductionHandler) ProductionHandler() paho.MessageHandler {
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

			ctx, cancel := context.WithTimeout(
				context.Background(),
				5*time.Second,
			)
			defer cancel()

			log.Println("===============Device 2 ===================")
			log.Println("===Device 2 ====>", req)
			log.Println("============================================")
			// ctx, cancel := context.WithTimeout(
			// 	context.Background(),
			// 	5*time.Second,
			// )
			// defer cancel()

			if err := h.service.Save(ctx, &req); err != nil {
				log.Printf("Save Error : %v", err)
				return
			}

			log.Printf(
				"Production Saved | EventID=%s | Device=%s | Station=%s | Count=%d | Cycle Time=%.2f sec",
				req.EventID,
				req.DeviceID,
				req.Station,
				req.Count,
				req.CycleTimeSec,
			)
		}(payload)
	}
}

// package handler

// import (
// 	"github.com/rajeshbond/smart/internal/mqtt/assembly/service"
// )

// type ProductionHandler struct {
// 	service service.ProductionSerive
// }

// func NewProductionHandler(service *service.ProductionSerive) *ProductionHandler {
// 	return &ProductionHandler{service: *service}
// }

// func (h *ProductionHandler) ProductionHandler() paho.MessageHandler {

// 	return func(client paho.Client, msg paho.Message) {

// 		payload := append([]byte(nil), msg.Payload()...)

// 		go func(data []byte) {

// 			var req dto.ProductionDTO

// 			if err := json.Unmarshal(data, &req); err != nil {

// 				log.Printf("Production JSON Error : %v", err)
// 				log.Printf("Raw Payload: %q", string(msg.Payload()))

// 				return
// 			}

// 			ctx, cancel := context.WithTimeout(
// 				context.Background(),
// 				5*time.Second,
// 			)

// 			defer cancel()

// 			if err := h.service.Save(ctx, &req); err != nil {

// 				log.Printf("Save Error : %v", err)

// 				return
// 			}
// 			// fmt.Println(req)
// 			log.Printf(
// 				"Production Saved | EventID=%s | Device=%s | Station=%s | Count=%d | Cycle Time=%.2f sec",
// 				req.EventID,
// 				req.DeviceID,
// 				req.Station,
// 				req.Count,
// 				req.CycleTimeSec,
// 			)
// 		}(payload)

// 	}

// }

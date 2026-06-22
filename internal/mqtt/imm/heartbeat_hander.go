package imm

import (
	"context"
	"encoding/json"
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/rajeshbond/smart/database"
)

type Heartbeat struct {
	TenantID  int64  `json:"tenant_id"`
	MachineID string `json:"machine_id"`
}

type WatchdogService struct {
	DB *database.DB
}

func NewWatchdogService(s *database.DB) *WatchdogService {
	return &WatchdogService{DB: s}
}

func (w *WatchdogService) Handler() paho.MessageHandler {

	return func(client paho.Client, msg paho.Message) {

		var hb Heartbeat

		if err := json.Unmarshal(msg.Payload(), &hb); err != nil {
			log.Println("heartbeat error:", err)
			return
		}

		_, err := w.DB.PGX.Exec(
			context.Background(),
			`
			INSERT INTO machine_status (tenant_id, machine_id, last_seen, status)
			VALUES ($1,$2,NOW(),'ONLINE')
			ON CONFLICT (tenant_id, machine_id)
			DO UPDATE SET last_seen = NOW(), status = 'ONLINE'
			`,
			hb.TenantID,
			hb.MachineID,
		)

		if err != nil {
			log.Println(err)
		}
	}
}

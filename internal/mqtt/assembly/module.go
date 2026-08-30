package assembly

import (
	"database/sql"

	"github.com/rajeshbond/smart/internal/mqtt/assembly/handler"
	"github.com/rajeshbond/smart/internal/mqtt/assembly/service"
	"github.com/rajeshbond/smart/internal/mqtt/assembly/store"
)

// ============================================================
// MQTT HANDLERS
// ============================================================

type MQTTHandlers struct {
	Production *handler.ProductionHandler
	Heartbeat  *handler.HeartbeatHandler
}

// ============================================================
// ASSEMBLY MODULE
// ============================================================

type Module struct {
	Store store.ProductionStore

	Service service.ProductionSerive

	Handlers MQTTHandlers
}

// ============================================================
// NEW MODULE
// ============================================================

func NewModule(db *sql.DB) *Module {

	// --------------------------------------------------------
	// Store
	// --------------------------------------------------------

	productionStore :=
		store.NewProductionStore(db)

	// --------------------------------------------------------
	// Production service
	// --------------------------------------------------------

	productionService :=
		service.NewProductionService(
			productionStore,
		)

	// --------------------------------------------------------
	// Production handler
	// --------------------------------------------------------

	productionHandler :=
		handler.NewProductionHandler(
			&productionService,
		)

	// --------------------------------------------------------
	// Heartbeat handler
	// --------------------------------------------------------

	heartbeatHandler :=
		handler.NewHeartbeatHandler()

	// --------------------------------------------------------
	// Module
	// --------------------------------------------------------

	return &Module{

		Store: productionStore,

		Service: productionService,

		Handlers: MQTTHandlers{

			Production: productionHandler,

			Heartbeat: heartbeatHandler,
		},
	}
}

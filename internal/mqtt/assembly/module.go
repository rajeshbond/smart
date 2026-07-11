package assembly

import (
	"database/sql"

	"github.com/rajeshbond/smart/internal/mqtt/assembly/handler"
	"github.com/rajeshbond/smart/internal/mqtt/assembly/service"
	"github.com/rajeshbond/smart/internal/mqtt/assembly/store"
)

type Module struct {
	Store   store.ProductionStore
	Service service.ProductionSerive
	Handler *handler.ProductionHandler
}

func NewModule(db *sql.DB) *Module {

	productionStore := store.NewProductionStore(db)

	productionService := service.NewProductionService(productionStore)

	productionHandler := handler.NewProductionHandler(&productionService)

	return &Module{
		Store:   productionStore,
		Service: productionService,
		Handler: productionHandler,
	}
}

package devicemaster

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/jmoiron/sqlx"
	"github.com/rajeshbond/smart/config"
	"github.com/rajeshbond/smart/internal/http/device/device_master/handler"
	"github.com/rajeshbond/smart/internal/http/device/device_master/service"
	"github.com/rajeshbond/smart/internal/http/device/device_master/store"
	"github.com/rajeshbond/smart/internal/mqtt/mqttadmin"
)

type Module struct {
	Handler   *handler.Handler
	Service   *service.Service
	Store     *store.Store
	MQTTAdmin mqttadmin.Service
	tokenAuth *jwtauth.JWTAuth
}

func NewModule(db *sqlx.DB, cfg *config.Config, tokenAuth *jwtauth.JWTAuth) *Module {
	//----------------------------------------------------------
	// Store
	//----------------------------------------------------------

	deviceStore := store.NewStore(db)
	//----------------------------------------------------------
	// MQTT Admin
	//----------------------------------------------------------
	mqttAdmin := mqttadmin.NewService(cfg)
	deviceService := service.NewService(deviceStore, mqttAdmin)

	//----------------------------------------------------------
	// Handler
	//----------------------------------------------------------

	deviceHandler := handler.NewHandler(

		deviceService,

		tokenAuth,
	)

	//----------------------------------------------------------
	// Module
	//----------------------------------------------------------

	return &Module{
		tokenAuth: tokenAuth,
		Handler:   deviceHandler,
		Service:   deviceService,
		Store:     deviceStore,
	}

}

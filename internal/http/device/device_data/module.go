package devicedata

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rajeshbond/smart/internal/http/device/device_data/handler"
	"github.com/rajeshbond/smart/internal/http/device/device_data/service"
	"github.com/rajeshbond/smart/internal/http/device/device_data/store"
)

type Module struct {
	Handler   *handler.Handler
	Service   *service.Service
	Store     *store.Store
	tokenAuth *jwtauth.JWTAuth
}

func NewModule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *Module {
	store := store.NewStore(db)
	service := service.NewService(store)
	handler := handler.NewService(service, tokenAuth)

	return &Module{
		tokenAuth: tokenAuth,
		Handler:   handler,
		Service:   service,
		Store:     store,
	}
}

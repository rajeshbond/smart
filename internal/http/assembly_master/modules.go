package assemblymaster

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rajeshbond/smart/internal/http/assembly_master/handler"
	"github.com/rajeshbond/smart/internal/http/assembly_master/service"
	"github.com/rajeshbond/smart/internal/http/assembly_master/store"
)

type Module struct {
	Handler   *handler.Handler
	Service   *service.Service
	Store     *store.Store
	tokenAuth *jwtauth.JWTAuth
}

func NewModuleAssembly(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *Module {
	store := store.NewStore(db)
	service := service.NewService(store)
	handler := handler.NewHandler(service, tokenAuth)

	return &Module{
		tokenAuth: tokenAuth,
		Handler:   handler,
		Service:   service,
		Store:     store,
	}

}

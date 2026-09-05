package mold

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	moldhandler "github.com/rajeshbond/smart/internal/http/mold/mold_handler"
	moldservice "github.com/rajeshbond/smart/internal/http/mold/mold_service"
	moldstore "github.com/rajeshbond/smart/internal/http/mold/mold_store"
)

type MoldModule struct {
	MoldHandler *moldhandler.MoldHanlder
	MoldService *moldservice.MoldSerice
	MoldStore   *moldstore.MoldStore
	tokenAuth   *jwtauth.JWTAuth
}

func NewMoldModeule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *MoldModule {
	moldStore := moldstore.NewMoldStore(db)
	moldService := moldservice.NewMoldService(moldStore)
	moldHandler := moldhandler.NewMoldHandler(moldService, tokenAuth)

	return &MoldModule{
		tokenAuth:   tokenAuth,
		MoldStore:   moldStore,
		MoldService: moldService,
		MoldHandler: moldHandler,
	}
}

package userrole

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
)

type Module struct {
	Handler   *Handler // Capitalized
	Service   *Service // Capitalized
	Store     *Store
	tokenAuth *jwtauth.JWTAuth // Capitalized
}

func NewModule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *Module {
	store := NewStore(db)
	service := NewService(store)
	handler := NewHandler(service, tokenAuth)

	return &Module{
		Handler: handler, // Use uppercase
		Service: service, // Use uppercase
		Store:   store,   // Use uppercase
	}
}

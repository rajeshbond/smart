package imm

import (
	"github.com/rajeshbond/smart/database"
)

type Module struct {
	Store   *Store
	Service *Service
	Handler *Handler
}

func NewModule(db *database.DB) *Module {
	store := NewStore(db.PGX)
	service := NewService(store)
	handler := NewHandler(service)

	return &Module{
		Store:   store,
		Service: service,
		Handler: handler,
	}
}

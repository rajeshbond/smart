package command

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (m *Module) Router() chi.Router {

	log.Println("Assembly Command Router Loaded")

	r := chi.NewRouter()

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Assembly Module Ok"))
	})

	r.Group(func(r chi.Router) {
		// r.Use(auth.Verifier(m.tokenAuth))
		// r.Use(auth.Authenticator(m.tokenAuth))
		// r.Use(auth.UserContextInjector)

		r.Post("/rest-counter", m.Handler.ResetCounterHandler)
	})

	return r
}

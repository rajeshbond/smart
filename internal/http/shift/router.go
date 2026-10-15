package shift

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rajeshbond/smart/internal/auth"
)

func (m *Module) Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Resource type Test Ok"))
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.Verifier(m.tokenAuth))
		r.Use(auth.Authenticator(m.tokenAuth))
		r.Use(auth.UserContextInjector)
		r.Post("/createshift", m.Handler.Create)

	})

	log.Println("========== ALL REGISTERED ROUTES ==========")

	if err := chi.Walk(r, func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler,
	) error {
		log.Printf("HTTP ROUTE: %-6s %s", method, route)
		return nil
	}); err != nil {
		log.Printf("Route walk error: %v", err)
	}

	log.Println("============================================")

	return r
}

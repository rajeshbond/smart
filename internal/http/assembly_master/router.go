package assemblymaster

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rajeshbond/smart/internal/auth"
)

func (m *Module) Router() chi.Router {
	r := chi.NewRouter()

	// log.Println("========== ASSEMBLY MASTER ROUTER CREATED ==========")

	r.Get("/test1", func(w http.ResponseWriter, req *http.Request) {
		log.Println("========== ASSEMBLY MASTER TEST HIT ==========")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Assembly master Test Ok"))
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.Verifier(m.tokenAuth))
		r.Use(auth.Authenticator(m.tokenAuth))
		r.Use(auth.UserContextInjector)

		r.Post("/createassembly", m.Handler.CreateAssembly)
	})

	if err := chi.Walk(r, func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler,
	) error {
		log.Printf("ASSEMBLY MODULE ROUTE: %-6s %s", method, route)
		return nil
	}); err != nil {
		log.Printf("Assembly module route error: %v", err)
	}

	return r
}

package devicemaster

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rajeshbond/smart/internal/auth"
)

func (m *Module) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/device", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Device Route  Test is Ok"))
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.Verifier(m.tokenAuth))
		r.Use(auth.Authenticator(m.tokenAuth))
		r.Use(auth.UserContextInjector)
		// Device CRUD
		r.Post("/", m.Handler.Create)
		r.Get("/", m.Handler.List)
		r.Put("/{deviceID}", m.Handler.Update)
		r.Delete("/{deviceID}", m.Handler.Delete)

		// MQTT Registration
		r.Post(
			"/{deviceID}/mqtt/register",
			m.Handler.RegisterMQTTUsername,
		)

		// MQTT Unregister (Future)
		r.Delete(
			"/{deviceID}/mqtt/register",
			http.NotFound,
		)
	})

	return r
}

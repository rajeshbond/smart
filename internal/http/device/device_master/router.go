package devicemaster

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rajeshbond/smart/internal/auth"
)

func (m *Module) Routes() chi.Router {
	r := chi.NewRouter()

	// Authentication
	r.Use(auth.Authenticator(m.tokenAuth))

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

	return r
}

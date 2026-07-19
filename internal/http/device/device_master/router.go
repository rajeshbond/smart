/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : router.go
 *
 * DESCRIPTION :
 * Device Master Routes
 *
 ******************************************************************************/

package devicemaster

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rajeshbond/smart/internal/auth"
)

func (m *Module) RegisterRoutes(r chi.Router) {

	r.Route("/", func(r chi.Router) {

		//----------------------------------------------------------------------
		// Authentication
		//----------------------------------------------------------------------

		r.Use(auth.Authenticator(m.tokenAuth))

		//----------------------------------------------------------------------
		// Device CRUD
		//----------------------------------------------------------------------

		r.Post("/", m.Handler.Create)

		r.Get("/", m.Handler.List)

		// r.Get("/{deviceID}", m.Handler.Get)

		r.Put("/{deviceID}", m.Handler.Update)

		r.Delete("/{deviceID}", m.Handler.Delete)

		//----------------------------------------------------------------------
		// Device Status
		//----------------------------------------------------------------------

		// r.Patch("/{deviceID}/status", m.Handler.UpdateStatus)

		//----------------------------------------------------------------------
		// Firmware
		//----------------------------------------------------------------------

		// r.Patch("/{deviceID}/firmware", m.Handler.UpdateFirmware)

		//----------------------------------------------------------------------
		// MQTT Registration
		//----------------------------------------------------------------------

		r.Post(
			"/{deviceID}/mqtt/register",
			m.Handler.RegisterMQTTUsername,
		)

		//----------------------------------------------------------------------
		// MQTT Unregister (Future)
		//----------------------------------------------------------------------

		r.Delete(
			"/{deviceID}/mqtt/register",
			http.NotFound,
		)
	})
}

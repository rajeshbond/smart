package application

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rajeshbond/smart/internal/http/command"
)

func NewRouter(app *App) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // React/Next.js frontend
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{"*"},
		// AllowedHeaders: []string{
		// 	"Accept",
		// 	"Authorization",
		// 	"Content-Type",
		// 	"X-CSRF-Token",
		// },
		// ExposedHeaders:   []string{"Link"},
		// AllowCredentials: true,
		// MaxAge:           300,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	commandModule := command.NewModule(app.MQTTClient)
	r.Mount("/api/v1/assembly-command", commandModule.Router())

	return r
}

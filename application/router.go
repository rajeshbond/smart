package application

//////////////////////////////////////////////
// Packages Mounted
//////////////////////////////////////////////
// 1.User Role
// 2.Tenant

// =======Command for MQTT Clients (esp 32)========
// MQTT Esp 32 reset count
// ================================================
//////////////////////////////////////////////
import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rajeshbond/smart/cmd/service"
	"github.com/rajeshbond/smart/internal/http/command"
	"github.com/rajeshbond/smart/internal/http/tenant"
	userrole "github.com/rajeshbond/smart/internal/http/user_role"
	"github.com/rajeshbond/smart/internal/http/users"
	httpSwagger "github.com/swaggo/http-swagger"
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

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health Ok"))
	})

	// Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	tokenAuth := service.GetTokenAuth()

	// ================================================================
	// Mouldes mounting (Starts)
	// ================================================================

	// 1.User Role---->

	userRoleModule := userrole.NewModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/user-role", userRoleModule.Router())

	// 2.Tenant---->

	tenantModule := tenant.NewModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/tenant", tenantModule.Router())

	//3. Users
	usersModule := users.NewModule(app.DB.SQLDB, tokenAuth, userRoleModule.Service, tenantModule.Service)
	r.Mount("/users", usersModule.Router())

	// Route Mounts Ends here <---------------

	// MQTT Commands ----->(MQTT)

	commandModule := command.NewModule(app.MQTTClient)
	r.Mount("/api/v1/assembly-command", commandModule.Router())
	// ================================================================
	// Mouldes mounting (Ends)
	// ================================================================

	return r
}

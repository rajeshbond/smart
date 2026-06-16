package application

import (
	"log"
	"net/http"

	"github.com/rajeshbond/smart/config"
	"github.com/rajeshbond/smart/database"
)

type App struct {
	DB     *database.DB
	Config *config.Config
}

func NewApp() *App {
	cfg := config.Load()
	db := database.NewDB(cfg)

	return &App{
		DB:     db,
		Config: cfg,
	}
}

func (a *App) Start() error {
	defer func() {
		if a.DB != nil {
			a.DB.Close()
		}
	}()

	r := NewRouter(a)

	log.Println("🚀 Server running on:", a.Config.APPPORT)

	return http.ListenAndServe(":"+a.Config.APPPORT, r)

}

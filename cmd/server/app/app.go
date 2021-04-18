package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"notes-service/internal/database"
	"notes-service/internal/handler"
	"notes-service/internal/logger"
)

type App struct {
	logger *logger.Logger
	router *mux.Router
}

func NewApp(logger *logger.Logger, router *mux.Router) *App {
	return &App{
		logger: logger,
		router: router,
	}
}

func (a *App) DefineDatabase() (*gorm.DB, error) {

	a.logger.OutputInfo("Init database...")

	initDatabase, err := database.InitDatabase()
	if err != nil {
		a.logger.OutputError(fmt.Sprintf("Could not init database %s", err))
		return nil, err
	}

	a.logger.OutputInfo("Migrate database...")

	err = database.MigrateDB(initDatabase)
	if err != nil {
		a.logger.OutputError(fmt.Sprintf("Could not migrate database %s", err))
		return nil, err
	}
	return initDatabase, nil
}

func (a *App) DefineRoutingAndServe(h *handler.Handler) {

	a.logger.OutputInfo("Define Routing...")

	a.router.Use(a.logRequest)
	a.router.HandleFunc("/api/note", h.PostNote)

	a.logger.OutputInfo("Started up successfully!")

	_ = http.ListenAndServe(":8080", a.router)
}

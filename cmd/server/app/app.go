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

func (a *App) DefineRouting(h *handler.Handler) {

	a.logger.OutputInfo("Define Routing...")

	a.router.Use(a.logRequest)
	a.router.HandleFunc("/api/note", h.PostNote).Methods(http.MethodPost)

	a.router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	})

	a.logger.OutputInfo("Started up successfully!")
}

func (a *App) Start() {
	_ = http.ListenAndServe(":8080", a.router)
}

func (a *App) GetRouter() *mux.Router {
	return a.router
}

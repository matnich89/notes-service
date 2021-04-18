package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"notes-service/cmd/server/app"
	"notes-service/internal/handler"
	"notes-service/internal/note"

	"notes-service/internal/logger"
)

func main() {
	globalLogger := logger.NewLogger()
	app := app.NewApp(globalLogger, mux.NewRouter())
	database, err := app.DefineDatabase()
	if err != nil {
		fmt.Println("Could not start")
	}
	noteHandler := handler.NewHandler(note.NewService(database), globalLogger)
	app.DefineRouting(noteHandler)

}

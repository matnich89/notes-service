package main

import (
	"fmt"
	"notes-service/internal/database"
)

type App struct {
}

func (app *App) Run() error {
	db, err := database.InitDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Starting up notes service....")
	app := &App{}
	if err := app.Run(); err != nil {
		panic(err)
	}
	fmt.Println("Started up successfully :)")
}

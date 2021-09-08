package main

import (
	"fmt"
	"net/http"

	"github.com/daparadoks/go-rest-api/internal/comment"
	"github.com/daparadoks/go-rest-api/internal/database"
	transportHTTP "github.com/daparadoks/go-rest-api/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting Up Our APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	app := App{
		Name:    "Comment API",
		Version: "1.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting up our REST API")
		log.Fatal(err)
	}
}

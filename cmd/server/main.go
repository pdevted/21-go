package main

import (
	"fmt"
	"net/http"

	"github.com/pdevted/go-rest-api/internal/comment"
	"github.com/pdevted/go-rest-api/internal/database"
	transportHTTP "github.com/pdevted/go-rest-api/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

// App - contains application information
// to database conencdtions
type App struct {
	Name    string
	Version string
}

// Run - sets up application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"Appversion": app.Version,
		}).Info("setting up application")

	fmt.Println("Setting Up APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		log.Error("failed to setup database")
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	log.Info("App startup successful")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{
		Name:    "Commenting Service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting up our REST API")
		log.Fatal(err)
	}
}

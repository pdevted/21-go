package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/pdevted/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database conencdtions
type App struct{}

// Run - sets up application
func (app *App) Run() error {
	fmt.Println("Setting Up APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}

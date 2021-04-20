package main

import "fmt"

// App - the struct which contains things like pointers
// to database conencdtions
type App struct{}

// Run - sets up application
func (app *App) Run() error {
	fmt.Println("Setting Up APP")
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

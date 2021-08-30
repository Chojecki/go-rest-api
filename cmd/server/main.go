package main

import "fmt"

// App - the struct which contains things like
// pointers to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Run the app")
	return nil
}

func main() {
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting the app")
		fmt.Println(err)
	}
}

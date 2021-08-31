package main

import (
	"fmt"
	"net/http"

	"github.com/Chojecki/go-rest-api/internal/comment"
	"github.com/Chojecki/go-rest-api/internal/database"
	transportHTTP "github.com/Chojecki/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	var err error
	db, err := database.NewDataBase()
	if err != nil {
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
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting the app")
		fmt.Println(err)
	}
}

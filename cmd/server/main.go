package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/islamghany/prod-rest/internals/transport/http"
)

// the struct where things likes pointers to the database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("running")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		return err
	}
	return nil

}
func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("something went wrong")
		fmt.Print(err)
	}

}

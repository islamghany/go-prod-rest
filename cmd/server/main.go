package main

import (
	"fmt"
	"github.com/islamghany/go-prod-rest/internals/comment"
	"github.com/islamghany/go-prod-rest/internals/database"
	transportHTTP "github.com/islamghany/go-prod-rest/internals/transport/http"
	"net/http"
)

// the struct where things likes pointers to the database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("running")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	// err = database.MigrateDB(db)

	// if err != nil {
	// 	fmt.Println("error")
	// 	return err
	// }

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
	fmt.Println("something went wrong")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("something went wrong")
		fmt.Print(err)
	}

}

// export DB_HOST=localhost
// export DB_PORT=5432
// export DB_USERNAME=postgres
// export DB_NAME=amigoscode
// export DB_PASSWORD=islamghany

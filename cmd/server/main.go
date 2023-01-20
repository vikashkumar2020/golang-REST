package main

import (
	"fmt"
	transportHTTP "golang/internal/transport/http"
	"net/http"
)

// App  - the struct which contains things like pointers
// to database connections

type App struct{

}

func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	err := http.ListenAndServe(":8080",handler.Router);
	if err !=nil{
		fmt.Print("failed to setup server")
		return err
	}
	return nil
}

func main(){
	fmt.Println("Go rest API")
	app :=App{}
	if err := app.Run();
	err!=nil{
		fmt.Println("Error starting up our Rest API")
		fmt.Println(err)
	}
}
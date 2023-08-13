package main

import (
	"log"
	"menunda-app/controllers"
	"os"

	"github.com/satriaprayoga/menunda"
)

func initMenunda() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	men := &menunda.Menunda{}
	err = men.New(path)
	if err != nil {
		log.Fatal(err)
	}
	men.AppName = "MyApp"

	men.InfoLog.Println("Debug is set to", men.Debug)

	handlers := &controllers.Controllers{
		App: men,
	}

	app := &application{
		App:  men,
		Ctrl: handlers,
	}

	app.App.Routes = app.routes()

	return app
}

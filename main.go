package main

import (
	"menunda-app/controllers"

	"github.com/satriaprayoga/menunda"
)

type application struct {
	App  *menunda.Menunda
	Ctrl *controllers.Controllers
}

func main() {
	c := initMenunda()
	c.App.ListenAndServe()
}

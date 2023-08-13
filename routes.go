package main

import (
	"github.com/labstack/echo/v4"
)

func (a *application) routes() *echo.Echo {
	//middleware

	//add routes

	//static
	a.App.Routes.Static("/public", "public")

	return a.App.Routes
}

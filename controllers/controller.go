package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/menunda"
)

type Controllers struct {
	App *menunda.Menunda
}

func (c *Controllers) Home(e echo.Context) error {
	return e.JSON(http.StatusOK, "home")
}

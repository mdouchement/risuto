package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mdouchement/risuto/config"
)

// Version returns the current version of this application.
func Version(c echo.Context) error {
	c.Set("handler_method", "Version")

	return c.JSON(http.StatusOK, echo.Map{
		"version": config.Cfg.Version,
	})
}

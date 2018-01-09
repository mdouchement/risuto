package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// IndexHome renders application home page.
func IndexHome(c echo.Context) error {
	c.Set("handler_method", "IndexHome")

	return c.Render(http.StatusOK, "index.home.tmpl", echo.Map{})
}

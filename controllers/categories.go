package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mdouchement/risuto/models"
)

type Categories struct{}

func NewCategories() *Categories {
	return &Categories{}
}

// List implements REST inteface.
func (cs *Categories) List(c echo.Context) error {
	c.Set("handler_method", "List")

	return c.JSON(http.StatusOK, models.GetAllCategories())
}

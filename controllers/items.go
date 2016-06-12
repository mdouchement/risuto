package controllers

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/mdouchement/risuto/models"
)

// Items handles wishlist items
type Items struct{}

// ItemsCreateParams are the params for Create methods
type ItemsCreateParams struct {
	ID           string   `structs:"id"`
	Name         string   `json:"name" structs:"name" binding:"required"`
	Descriptions []string `json:"descriptions" structs:"descriptions" binding:"required"`
}

// ItemsUpdateParams are the params for Create methods
type ItemsUpdateParams struct {
	ID           string   `structs:"id" binding:"required"`
	Name         string   `json:"name" structs:"name"`
	Descriptions []string `json:"descriptions" structs:"descriptions"`
	Score        int      `json:"score" structs:"score"`
}

// NewItems creates a new Items controller
func NewItems() *Items {
	return &Items{}
}

// Create appends item
func (i *Items) Create(c *gin.Context) {
	var icp ItemsCreateParams
	if err := c.BindJSON(&icp); err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(fmt.Errorf("CreateItems BindJSON %v", err)))
		return
	}
	params := structs.New(icp).Map()

	item, err := models.CreateItem(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(fmt.Errorf("CreateItems: %v", err)))
		return
	}

	c.JSON(http.StatusCreated, item)
}

// Update edits the given item
func (i *Items) Update(c *gin.Context) {
	var iup ItemsUpdateParams
	if err := c.BindJSON(&iup); err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(fmt.Errorf("UpdateItem BindJSON %v", err)))
		return
	}
	params := structs.New(iup).Map()

	item, err := models.UpdateItem(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(fmt.Errorf("UpdateItem: %v", err)))
		return
	}

	c.JSON(http.StatusOK, item)
}

// List returns all items of the wishlist
func (i *Items) List(c *gin.Context) {
	c.JSON(http.StatusOK, models.AllItems())
}

// Delete removes the given item
func (i *Items) Delete(c *gin.Context) {
	if err := models.DestroyItem(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(fmt.Errorf("DeleteItem: %v", err)))
		return
	}

	c.Status(http.StatusNoContent)
}

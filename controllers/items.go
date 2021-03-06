package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mdouchement/risuto/models"
)

type (
	Items      []*models.Item
	itemParams struct {
		Name         string   `json:"name" valid:"required~required"`
		Category     string   `json:"category" valid:"required~required"`
		Descriptions []string `json:"descriptions" structs:"descriptions"`
		Score        int      `json:"score" structs:"score"`
	}
)

func (p *itemParams) RName() string {
	return "item"
}

func NewItems() *Items {
	return &Items{}
}

func (is *Items) Create(c echo.Context) error {
	c.Set("handler_method", "Create")

	// Filter params
	var params itemParams
	if err := c.Bind(&params); err != nil {
		return err
	}

	if err := c.Validate(&params); err != nil {
		return err
	}

	// Append params to the model
	item := models.NewItem()
	if err := MergeParams(item, params); err != nil {
		return err
	}

	// Persist the model
	if err := item.Save(); err != nil {
		return err
	}

	// Synchronize category collection and items' index
	models.IndexItem(item)

	return c.JSON(http.StatusCreated, item)
}

// List implements REST inteface.
func (is *Items) List(c echo.Context) error {
	c.Set("handler_method", "List")

	return c.JSON(http.StatusOK, models.GetAllFilteredItems(c.QueryParam("category")))
}

// Update implements REST inteface.
func (is *Items) Update(c echo.Context) error {
	c.Set("handler_method", "Update")

	// Filter params
	var params itemParams
	if err := c.Bind(&params); err != nil {
		return err
	}

	if err := c.Validate(&params); err != nil {
		return err
	}

	// Retrieve stored item
	item, err := models.GetItem(c.Param("id")) // FIXME Handle 404 not found well
	if err != nil {
		return err
	}

	// Synchronize category collection and items' index
	updateIndex := params.Category != item.Category
	if updateIndex {
		models.DeindexItem(item)
	}

	// Update attributes
	if err := MergeParams(item, params); err != nil {
		return err
	}
	if err := item.Save(); err != nil {
		return err
	}

	// Synchronize category collection
	if updateIndex {
		models.IndexItem(item)
	}

	return c.JSON(http.StatusOK, item)
}

// Delete implements REST inteface.
func (is *Items) Delete(c echo.Context) error {
	c.Set("handler_method", "Delete")

	item, err := models.GetItem(c.Param("id")) // FIXME Handle 404 not found well
	if err != nil {
		return err
	}

	if err := item.Delete(); err != nil {
		return err
	}

	// Synchronize category collection and items' index
	models.DeindexItem(item)

	return c.NoContent(http.StatusNoContent)
}

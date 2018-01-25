package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mdouchement/risuto/models"
)

type (
	Items      []*models.Item
	itemParams struct {
		ResourceName string   `json:"-"` // Used for error formating
		Name         string   `json:"name" valid:"required~required"`
		Category     string   `json:"category" valid:"required~required"`
		Descriptions []string `json:"descriptions" structs:"descriptions"`
		Score        int      `json:"score" structs:"score"`
	}
)

func (p *itemParams) RName() string {
	return p.ResourceName
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
	params.ResourceName = "item" // Needed for validation

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

	// Synchronize category collection
	models.IncrementCategory(item.Category)

	return c.JSON(http.StatusCreated, item)
}

// List implements REST inteface.
func (is *Items) List(c echo.Context) error {
	c.Set("handler_method", "List")

	return c.JSON(http.StatusOK, models.GetAllItems())
}

// Update implements REST inteface.
func (is *Items) Update(c echo.Context) error {
	c.Set("handler_method", "Update")

	// Filter params
	var params itemParams
	if err := c.Bind(&params); err != nil {
		return err
	}
	params.ResourceName = "item" // Needed for validation

	if err := c.Validate(&params); err != nil {
		return err
	}

	// Retrieve stored item
	item, err := models.GetItem(c.Param("id")) // FIXME Handle 404 not found well
	if err != nil {
		return err
	}

	// Synchronize category collection
	if params.Category != item.Category {
		models.IncrementCategory(params.Category)
		models.DecrementCategory(item.Category)
	}

	// Update attributes
	if err := MergeParams(item, params); err != nil {
		return err
	}
	if err := item.Save(); err != nil {
		return err
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

	// Synchronize category collection
	models.DecrementCategory(item.Category)

	return c.NoContent(http.StatusNoContent)
}

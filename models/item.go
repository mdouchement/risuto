package models

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"
	"github.com/mdouchement/risuto/config"
	"github.com/mdouchement/risuto/errors"
	"github.com/mdouchement/risuto/util"
)

// Item model.
type Item struct {
	ID           string   `json:"id" structs:"id"`
	Name         string   `json:"name" structs:"name"`
	Descriptions []string `json:"descriptions" structs:"descriptions"`
	Score        int      `json:"score" structs:"score"`
}

var itemCol = config.DB.Use("items")

// GetAllItems returns all existing entry of Dataset model from the database.
func GetAllItems() []*Item {
	items := []*Item{}
	itemCol.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
		var item Item
		json.Unmarshal(docContent, &item)
		item.ID = fmt.Sprintf("%d", id)
		items = append(items, &item)
		return true
	})
	return items
}

// GetItem returns a new instance of Item model from the database for the given id.
func GetItem(id string) (*Item, error) {
	doc, err := itemCol.Read(util.MustAtoi(id))
	if err != nil {
		return nil, errors.NewModelsError("not_found", errors.M{
			"resource": "item",
			"id":       id,
			"internal": err.Error(),
		})
	}

	data, err := json.Marshal(doc)
	if err != nil {
		return nil, errors.NewModelsError("unexpected", errors.M{
			"reason":  "Could not blobize found item",
			"action":  "marshalize",
			"details": err.Error(),
		})
	}

	var item Item
	if err = json.Unmarshal(data, &item); err != nil {
		return nil, errors.NewModelsError("unexpected", errors.M{
			"reason":  "Could not blobize found item",
			"action":  "unmarshalize",
			"details": err.Error(),
		})
	}
	return &item, nil
}

// NewItem returns a new instance of Item model.
func NewItem() *Item {
	return &Item{
		ID: "new-item",
	}
}

// Save inserts or updates the entry in database with its own fields.
func (m *Item) Save() error {
	if m.ID == "new-item" {
		id, err := itemCol.Insert(structs.Map(m))
		if err != nil {
			return errors.NewModelsError("unexpected", errors.M{
				"reason":  "Could not create item",
				"details": err.Error(),
			})
		}
		m.ID = fmt.Sprintf("%d", id)
	}

	if err := itemCol.Update(util.MustAtoi(m.ID), structs.Map(m)); err != nil {
		return errors.NewModelsError("unexpected", errors.M{
			"reason":  "Could not persist the item",
			"details": err.Error(),
		})
	}
	return nil
}

// Delete removes the entry in database.
func (m *Item) Delete() error {
	if err := itemCol.Delete(util.MustAtoi(m.ID)); err != nil {
		return errors.NewModelsError("unexpected", errors.M{
			"reason":  "Could not remove the item",
			"details": err.Error(),
		})
	}
	return nil
}

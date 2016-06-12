package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mdouchement/risuto/config"
	"github.com/mdouchement/risuto/utils"
)

// Item model
type Item map[string]interface{}

var itemCol = config.DB.Use("items")

// CreateItem inserts a new Item in the database
func CreateItem(item Item) (Item, error) {
	item["score"] = 0

	id, err := itemCol.Insert(item)
	item["id"] = fmt.Sprintf("%d", id)
	return item, err
}

// UpdateItem inserts a new Item in the database
func UpdateItem(item Item) (Item, error) {
	idx, err := strconv.Atoi(utils.String(item["id"]))
	if err != nil {
		return nil, fmt.Errorf("UpdateItemModel: %v", err)
	}

	err = itemCol.Update(idx, item)
	return item, err
}

// AllItems returns all Items persisted in the database
func AllItems() []Item {
	items := []Item{}
	itemCol.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
		var item Item
		json.Unmarshal(docContent, &item)
		item["id"] = fmt.Sprintf("%d", id)
		items = append(items, item)
		return true
	})
	return items
}

// DestroyItem removes the given item form the database
func DestroyItem(id string) error {
	idx, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("DestroyItemModel: %v", err)
	}
	return itemCol.Delete(idx)
}

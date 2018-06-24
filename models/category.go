package models

import (
	"sync"
)

type categories struct {
	sync.Mutex
	index map[string]map[string]struct{}
}

var cats categories

func init() {
	cats = categories{
		index: map[string]map[string]struct{}{},
	}

	for _, item := range GetAllItems() {
		IndexItem(item)
	}
}

// GetAllCategories returns all existing entry of Category model from the database.
func GetAllCategories() []string {
	cats.Lock()
	defer cats.Unlock()

	cs := []string{}
	for c := range cats.index {
		cs = append(cs, c)
	}

	return cs
}

// IndexItem indexs the given item according to its category to the database.
func IndexItem(item *Item) {
	cats.Lock()
	defer cats.Unlock()

	if cats.index[item.Category] == nil {
		cats.index[item.Category] = map[string]struct{}{}
	}
	cats.index[item.Category][item.ID] = struct{}{}
}

// DeindexItem removes the given item from the index in the database.
func DeindexItem(item *Item) {
	cats.Lock()
	defer cats.Unlock()

	if cats.index[item.Category] == nil {
		return
	}
	delete(cats.index[item.Category], item.ID)

	if len(cats.index[item.Category]) == 0 {
		delete(cats.index, item.Category)
	}
}

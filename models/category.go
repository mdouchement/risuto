package models

import "sync"

type categories struct {
	sync.Mutex
	values map[string]int
}

var cats categories

func init() {
	cats = categories{
		values: map[string]int{},
	}
	for _, item := range GetAllItems() {
		cats.values[item.Category]++
	}
}

// GetAllCategories returns all existing entry of Category model from the database.
func GetAllCategories() []string {
	cats.Lock()
	defer cats.Unlock()

	cs := []string{}
	for c := range cats.values {
		cs = append(cs, c)
	}

	return cs
}

// IncrementCategory increments the given category to the database.
func IncrementCategory(c string) {
	cats.Lock()
	defer cats.Unlock()

	cats.values[c]++
}

// DecrementCategory decrements or removes the given category to the database.
func DecrementCategory(c string) {
	cats.Lock()
	defer cats.Unlock()

	cats.values[c]--
	if cats.values[c] == 0 {
		delete(cats.values, c)
	}
}

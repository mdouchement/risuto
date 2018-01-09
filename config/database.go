package config

import "github.com/HouzuoGuo/tiedot/db"

// DB is the database instance
var DB *db.DB

func init() {
	var err error
	DB, err = db.OpenDB(Cfg.DBDir)
	check(err)

	// Migration
	collections := DB.AllCols()
	if !contains(collections, "items") {
		err := DB.Create("items")
		check(err)
	}

	// Repair and compact
	DB.Scrub("items")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

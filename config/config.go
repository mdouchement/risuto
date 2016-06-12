package config

import (
	"os/exec"
	"path/filepath"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/caarlos0/env"
)

var (
	// Cfg contains all current configuration
	Cfg *Config
	// DB is the database instance
	DB *db.DB
)

// Config holds all configuration for our program
type Config struct {
	Version   string `env:"APP_VERSION"`
	Namespace string `env:"RISUTO_NAMESPACE"`
	DBDir     string `env:"RISUTO_DATABASE" envDefault:"tmp/database"`
}

func init() {
	Cfg = &Config{}
	env.Parse(Cfg)

	dir, err := filepath.Abs(Cfg.DBDir)
	check(err)
	Cfg.DBDir = dir

	if Cfg.Version == "" {
		Cfg.Version = version()
	}

	if Cfg.Namespace != "" {
		Cfg.Namespace = filepath.Join("/", Cfg.Namespace)
	}

	initDB()
}

func version() string {
	output, _ := exec.Command("git", "describe", "--tags", "2>", "/dev/null").Output()
	version := string(output)
	if version == "" {
		version = "0.0.1"
	}
	return version
}

func initDB() {
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

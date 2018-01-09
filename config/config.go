package config

import (
	"path/filepath"

	"github.com/caarlos0/env"
)

// Cfg variable holds all configuration.
var Cfg *Config

// Config holds all configuration for our program.
type Config struct {
	Version string
	DBDir   string `env:"RISUTO_DATABASE" envDefault:"tmp/database"`

	RouterNamespace string `env:"ROUTER_NAMESPACE" envDefault:""`
}

func init() {
	Cfg = &Config{}
	env.Parse(Cfg)

	dir, err := filepath.Abs(Cfg.DBDir)
	check(err)
	Cfg.DBDir = dir
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

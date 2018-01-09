package main

import (
	"fmt"
	"os"

	"github.com/mdouchement/risuto/config"
	"github.com/mdouchement/risuto/web"
	"gopkg.in/urfave/cli.v2"
)

var (
	version = "dev"
	app     *cli.App
)

func init() {
	config.Cfg.Version = version

	app = &cli.App{
		Name:    "risuto",
		Version: config.Cfg.Version,
		Commands: []*cli.Command{
			web.Command,
		},
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

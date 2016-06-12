package main

import (
	"os"

	"github.com/mdouchement/risuto/config"
	"github.com/mdouchement/risuto/web"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Risuto"
	app.Version = config.Cfg.Version
	app.Usage = "Wishlist webserver"
	// TODO add subcommands `server` and `database`
	app.Flags = flags
	app.Action = action

	err := app.Run(os.Args)
	if err != nil {
		println(err)
	}
}

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:  "p, port",
		Usage: "Specify the port to listen to.",
	},
	&cli.StringFlag{
		Name:  "b, binding",
		Usage: "Binds server to the specified IP.",
	},
}

func action(context *cli.Context) error {
	port := context.String("p")
	if port == "" {
		port = "5000"
	}
	return web.Server(context.String("b"), port)
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdouchement/risuto/config"
	"github.com/mdouchement/risuto/util"
	"github.com/mdouchement/risuto/web"
	"github.com/pkg/errors"
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
			assetsFetchCommand,
			web.Command,
		},
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		config.Log.Error(err)
	}
}

//-------------------------------------//
//                                     //
// Assets fetcher                      //
//                                     //
//-------------------------------------//

var (
	assetsFetchCommand = &cli.Command{
		Name:    "fetch",
		Aliases: []string{"f"},
		Usage:   "fetch assets",
		Action:  assetsFetchAction,
		Flags:   assetsFetchFlags,
	}

	assetsFetchFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "min",
			Usage: "Minified version.",
		},
	}
)

func assetsFetchAction(context *cli.Context) error {
	data, err := ioutil.ReadFile("assets-manifest.json")
	if err != nil {
		return errors.Wrap(err, "Could not load assets-manifest.json")
	}

	var manifest map[string]string
	if err = json.Unmarshal(data, &manifest); err != nil {
		return errors.Wrap(err, "Could not parse assets-manifest.json")
	}

	for filename, url := range manifest {
		if !context.Bool("min") && strings.Contains(url, "cdnjs.cloudflare.com") && strings.Contains(url, ".min.") {
			filename = strings.Replace(filename, ".min", "", -1)
			url = strings.Replace(url, ".min", "", -1)
		}

		config.Log.Infof("Downloading asset %s", url)
		dst, err := filepath.Abs(filepath.Join("public", filename))
		if err != nil {
			return errors.Wrap(err, "Could not get absolute path")
		}

		util.MkdirAllWithFilename(dst)
		if err = util.Download(url, dst); err != nil {
			return errors.Wrap(err, "Could not get assets")
		}
	}

	return nil
}

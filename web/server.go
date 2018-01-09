package web

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mdouchement/risuto/config"
	"github.com/mdouchement/risuto/controllers"
	"github.com/mdouchement/risuto/web/middlewares"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
	"gopkg.in/urfave/cli.v2"
)

var (
	// Command defines the server command (CLI).
	Command = &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "start server",
		Action:  action,
		Flags:   flags,
	}

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "p, port",
			Usage: "Specify the port to listen to.",
		},
		&cli.StringFlag{
			Name:  "b, binding",
			Usage: "Binds server to the specified IP.",
		},
	}

	assets = packr.NewBox("../public")
)

func action(context *cli.Context) error {

	engine := EchoEngine()
	printRoutes(engine)

	listen := fmt.Sprintf("%s:%s", context.String("b"), context.String("p"))
	config.Log.Infof("Server listening on %s", listen)
	engine.Start(listen)

	return nil
}

// EchoEngine instanciates the LSS server.
func EchoEngine() *echo.Echo {
	engine := echo.New()
	engine.Use(middleware.Recover())
	engine.Use(middlewares.Echorus(config.Log))
	// Error handler
	engine.HTTPErrorHandler = middlewares.HTTPErrorHandler
	// Views templates
	engine.Renderer = templates
	// Strong parameters
	engine.Validator = &controllers.ParamsValidator{}

	router := engine.Group(config.Cfg.RouterNamespace)

	// Assets
	router.Use(assetsFS("/public", assets))

	router.GET("/version", controllers.Version)
	router.GET("/", controllers.IndexHome)
	middlewares.CRUD(router, "/items", controllers.NewItems())

	return engine
}

func assetsFS(urlPrefix string, assets packr.Box) echo.MiddlewareFunc {
	var fs = http.FileServer(assets)
	var fsm = fs

	if config.Env() == config.Production {
		m := minify.New()
		m.AddFunc("text/css", css.Minify)
		m.AddFunc("text/html", html.Minify)
		m.AddFunc("application/javascript", minifyVueJS)
		m.AddFunc("application/x-javascript", minifyVueJS)
		m.AddFunc("image/svg+xml", svg.Minify)
		m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
		m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
		fsm = m.Middleware(fs)
	}

	return func(before echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := before(c)
			if err != nil {
				if c, ok := err.(*echo.HTTPError); !ok || c.Code != http.StatusNotFound {
					return err
				}
			}

			w, r := c.Response(), c.Request()

			p := strings.TrimPrefix(r.URL.Path, urlPrefix)
			if !w.Committed && assets.Has(p) {
				r.URL.Path = p
				if strings.Contains(p, ".min.") {
					fs.ServeHTTP(w, r)
				} else {
					fsm.ServeHTTP(w, r)
				}
				return nil
			}
			return nil
		}
	}
}

func printRoutes(e *echo.Echo) {
	fmt.Println("Routes:")
	for _, route := range e.Routes() {
		fmt.Printf("%6s %s\n", route.Method, route.Path)
	}
}

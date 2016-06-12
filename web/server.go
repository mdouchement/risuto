package web

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/mdouchement/risuto/config"
	"github.com/mdouchement/risuto/controllers"
	"github.com/mdouchement/risuto/web/middlewares"
)

// Server routes all requests to controllers
func Server(binding, port string) error {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middlewares.Ginrus(config.Log))
	engine.Use(middlewares.ParamsConverter())

	engine.LoadHTMLGlob("views/*")

	router := engine.Group(namespace("/"))
	router.Static("/public", "public")
	router.GET("/", controllers.IndexHome)
	router.GET("/version", controllers.ShowVersion)
	middlewares.CRUD(router, "/items", controllers.NewItems())

	listener := fmt.Sprintf("%s:%s", binding, port)
	config.Log.Infof("Server listening on %s", listener)
	engine.Run(listener)

	return nil
}

func namespace(route string) string {
	return filepath.Join(config.Cfg.Namespace, route)
}

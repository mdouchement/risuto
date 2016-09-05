package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/mdouchement/risuto/config"
)

// IndexHome renders Risuto veropn
func IndexHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.index.tmpl", gin.H{
		"namespace": config.Cfg.Namespace,
		// "base":      "https://unpkg.com",
		"base": filepath.Join(config.Cfg.Namespace, "/public/js/vendor"),
	})
}

package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigRouter(app *HexoEditAndDeploy) {
	app.Router.LoadHTMLGlob("templates/*")

	app.Router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	app.Router.GET("/createNew", func(c *gin.Context) {
		c.HTML(http.StatusOK, "newBlog.tmpl", nil)
	})
}

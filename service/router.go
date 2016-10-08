package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigRouter(app *HexoEditAndDeploy) {
	app.Router.LoadHTMLGlob("templates/*")

	app.Router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	app.Router.GET("/createNew", func(c *gin.Context) {
		c.HTML(http.StatusOK, "newBlog.html", nil)
	})

	app.Router.POST("/push", func(c *gin.Context) {
		//content := c.Param("content")

		c.HTML(http.StatusOK, "newBlog.html", nil)
	})
}

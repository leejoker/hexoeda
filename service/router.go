package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigRouter(app *HexoEditAndDeploy) {
	app.Router.LoadHTMLGlob("templates/*")

	app.Router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	app.Router.GET("/createNew", func(c *gin.Context) {
		c.HTML(http.StatusOK, "newBlog.html", nil)
	})

	app.Router.POST("/obtainContent", func(c *gin.Context) {
		content := c.PostForm("content")
		path := app.Conf.HexoPath
		hexo := NewHexo("", content, path)
		CreateNewBlog(hexo)
		Clean(hexo)
		Generate(hexo)
		Deploy(hexo)
		c.HTML(http.StatusOK, "", nil)
	})
}

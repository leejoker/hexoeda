package service

import (
	"fmt"
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
		fmt.Print("内容：\n" + content)

		//此处进行文本处理，生成对应的md文件并执行hexo发布
		path := app.Conf.HexoPath
		hexo := NewHexo("test", content, path)
		CreateNewBlog(hexo)
		Clean(hexo)
		Generate(hexo)
		StartServer(hexo)
		c.HTML(http.StatusOK, "", nil)
	})
}

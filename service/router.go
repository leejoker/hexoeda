package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"fmt"
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

	app.Router.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", nil)
	})

	app.Router.POST("/blogList", func(c *gin.Context) {
		path := app.Conf.HexoPath
		posts := path + "/source/_posts"
		//fmt.Println(posts)
		dir_list, e := ioutil.ReadDir(posts)
		if e != nil {
			fmt.Println(" read dir error")
			c.HTML(http.StatusOK, "error.html", gin.H{
				"msg": "Find Nothing!",
			})
		}else{
			filenames := ""
			for _, v := range dir_list {
				//fmt.Println(i, "=", v.Name())
				filenames = filenames + posts + "/" + v.Name() + ";"
			}
			filenames = filenames[0:len(filenames) - 1]
			c.JSON(http.StatusOK, gin.H{
				"files": filenames,
			})
		}
	})
}

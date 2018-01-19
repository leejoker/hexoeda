package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

	app.Router.POST("/findContent", func(c *gin.Context) {
		fileName := c.PostForm("name")
		if fileName == "" {
			fmt.Println("no parameter")
			c.HTML(http.StatusOK, "error.html", gin.H{
				"msg": "No Parameter!",
			})
		} else {
			path := app.Conf.HexoPath
			filePath := path + "/source/_posts/" + fileName
			fin, err := os.Open(filePath)
			defer fin.Close()
			if err != nil {
				fmt.Println(filePath, err)
				return
			}
			buff := make([]byte, 1024)

			var buf bytes.Buffer
			buf.WriteString("")

			for {
				n, _ := fin.Read(buff)
				if 0 == n {
					break
				}
				buf.WriteString(string(buff[0:n]))
			}

			fileContent := buf.String()

			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"error":  nil,
				"data":   fileContent,
			})
		}
	})

	app.Router.POST("/obtainContent", func(c *gin.Context) {
		content := c.PostForm("content")
		path := app.Conf.HexoPath
		hexo := NewHexo("", content, path)
		CreateNewBlog(hexo)
		Clean(hexo)
		Generate(hexo)
		Deploy(hexo)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"error":  nil,
		})
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
		} else {
			filenames := ""
			for _, v := range dir_list {
				//fmt.Println(i, "=", v.Name())
				filenames = filenames + posts + "/" + v.Name() + ";"
			}
			filenames = filenames[0 : len(filenames)-1]
			c.JSON(http.StatusOK, gin.H{
				"files": filenames,
			})
		}
	})
}

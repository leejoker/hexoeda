package service

import (
	"github.com/go-playground/log"
	"os"
	"os/exec"
)

type Hexo struct {
	Title   string `the blog title`
	Content string `the blog content`
	Path    string `the path of hexo directory`
}

func NewHexo(title string, content string, path string) *Hexo {
	return &Hexo{Title: title, Content: content, Path: path}
}

func CreateNewBlog(hexo *Hexo) (result bool, err error) {
	cmd := exec.Command("hexo", "--cwd", hexo.Path, "new", hexo.Title)
	err = cmd.Start()
	if err != nil {
		log.Info("create new blog faild")
		return false, err
	}

	//copy the content to md file
	filename := hexo.Path + "/" + hexo.Title + ".md"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		log.Info("file not found")
		return false, err
	}
	defer file.Close()
	_, err = file.WriteString(hexo.Content)
	if err != nil {
		log.Info("write file error")
		return false, err
	}
	return true, err
}

func Deploy(hexo *Hexo) {
	cmd := exec.Command("hexo", "--cwd", hexo.Path, "deploy")
	err := cmd.Start()
	if err != nil {
		log.Info(err)
	}
}

func Clean(hexo *Hexo) {
	cmd := exec.Command("hexo", "--cwd", hexo.Path, "clean")
	err := cmd.Start()
	if err != nil {
		log.Info(err)
	}
}

func Generate(hexo *Hexo) {
	cmd := exec.Command("hexo", "--cwd", hexo.Path, "generate")
	err := cmd.Start()
	if err != nil {
		log.Info(err)
	}
}

func StartServer(hexo *Hexo) {
	cmd := exec.Command("hexo", "--cwd", hexo.Path, "s")
	err := cmd.Start()
	if err != nil {
		log.Info(err)
	}
}

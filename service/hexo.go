package service

import (
	"bufio"
	"fmt"
	"github.com/go-playground/log"
	"io"
	"os"
	"os/exec"
	"strings"
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
	log.Info("hexo地址为：" + hexo.Path)
	//log.Info("hexo内容为：" + hexo.Content)
	//through the content to obtain the title
	cotentArray := strings.Split(hexo.Content, "\n")
	titleLine := cotentArray[1]
	titleArray := strings.Split(titleLine, ":")
	hexo.Title = strings.Replace(titleArray[1], " ", "", -1)
	hexo.Title = strings.Replace(hexo.Title, "\n", "", -1)
	hexo.Title = strings.Replace(hexo.Title, "\r", "", -1)
	//create hexo blog file
	filename := hexo.Path + "/source/_posts/" + hexo.Title + ".md"
	log.Info("the filename is :" + filename)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Info("blog file create faild")
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
	params := []string{"--cwd", hexo.Path, "deploy"}
	execCommand("hexo", params)
}

func Clean(hexo *Hexo) {
	cmd := exec.Command("hexo", "--cwd", hexo.Path, "clean")
	err := cmd.Start()
	if err != nil {
		log.Info(err)
	}
}

func Generate(hexo *Hexo) {
	params := []string{"--cwd", hexo.Path, "generate"}
	execCommand("hexo", params)
}

func StartServer(hexo *Hexo) {
	cmd := exec.Command("hexo", "--cwd", hexo.Path, "server")
	err := cmd.Start()
	if err != nil {
		log.Info(err)
	}
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Printf("执行命令: %s\n", strings.Join(cmd.Args[1:], " "))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
		return false
	}
	cmd.Start() // Start开始执行c包含的命令，但并不会等待该命令完成即返回。Wait方法会返回命令的返回状态码并在命令返回后释放相关的资源。

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}

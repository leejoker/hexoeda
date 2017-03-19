package main

import (
	"github.com/leejoker/hexoeda/service"
)

func main() {
	app, err := service.NewApplication("config.conf")
	if err != nil {
		panic(err)
	}
	app.Router.Run(":8080")
}

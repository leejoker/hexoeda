package service

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

var (
	LogLevels = []log.Level{
		log.InfoLevel,
		log.NoticeLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.PanicLevel,
		log.AlertLevel,
		log.FatalLevel,
	}
)

type HexoEditAndDeploy struct {
	Router *gin.Engine
	Conf   *Config
}

func NewApplication(configFile string) (*HexoEditAndDeploy, error) {
	conf := ReadConfigFile(configFile)

	//add log
	cLog := console.New()
	log.RegisterHandler(cLog, LogLevels...)
	gin.SetMode(gin.ReleaseMode)

	log.Info("init log")
	//gin settings
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(static.Serve("/static/", static.LocalFile("./static", true)))

	app := &HexoEditAndDeploy{Router: router, Conf: conf}
	log.Info("create app over")

	//set router
	ConfigRouter(app)
	log.Info("set router over")

	return app, nil
}

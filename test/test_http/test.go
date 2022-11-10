package test_http

import (
	"b5gocmf/router"
	"b5gocmf/utils/core"
	"github.com/gin-gonic/gin"
	"io"
)

func TestLoadEnv()  *gin.Engine{
	core.TestLoad("../../")
	gin.SetMode(gin.ReleaseMode)

	gin.DefaultWriter = io.Discard
	engine := gin.New()
	//使用自定义格式的日志log

	engine.Use(gin.LoggerWithFormatter(core.CustomB5Log)).Use(gin.Recovery())

	router.LoadRouter(engine)
	return engine
}

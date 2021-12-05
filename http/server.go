package http

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() error {
	engine := gin.New()
	addr := viper.GetString("http.server")
	return engine.Run(addr)
}

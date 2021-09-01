package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IController interface {
	Register(*gin.Engine)
}

func InitRouter() *gin.Engine {
	en := gin.Default()
	en.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong!")
	})
	return en
}

func Start(en *gin.Engine) {
	if err := en.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

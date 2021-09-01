package user

import (
	"crud-api/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	router.IController
	service Service
}

func InitController(service Service) Controller {
	return Controller{service: service}
}

func (ctrl *Controller) Register(en *gin.Engine) {
	userGroup := en.Group("/users")
	userGroup.GET("", ctrl.findAll)
	userGroup.GET(":id", ctrl.findById)
}

func (ctrl *Controller) findAll(ctx *gin.Context) {
	users := ctrl.service.FindAll()
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (ctrl *Controller) findById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := ctrl.service.FindById(uint(id))
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

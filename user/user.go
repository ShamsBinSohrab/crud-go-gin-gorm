package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Service struct {
	DB *gorm.DB
}

type Controller struct {
	service Service
}

func ProvideUserService(DB *gorm.DB) Service {
	return Service{DB: DB}
}

func (s *Service) FindAll() []User {
	var users []User
	s.DB.Find(&users)
	return users
}

func (s *Service) FindById(id uint) User {
	var user User
	s.DB.First(&user, id)
	return user
}

func ProvideUserController(service Service) Controller {
	return Controller{service: service}
}

func (c *Controller) FindAll(ctx *gin.Context)  {
	users := c.service.FindAll()
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *Controller) FindById(ctx *gin.Context)  {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	user := c.service.FindById(uint(id))
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

package main

import (
	"crud-api/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

const host = "localhost"
const username = "aiaibot"
const password = "lena"
const dbname = "testdb"
const port = 5432

var database Database

type Database struct {
	*gorm.DB
}

func (db *Database) doMigration() error {
	err := db.AutoMigrate(&user.User{})
	return err
}

func (db *Database) seedData() {
	db.Create(&user.User{
		Model: gorm.Model{},
		Name:  "Shams Bin Sohrab",
		Email: "aniknishan@gmail.com",
	})
	db.Create(&user.User{
		Model: gorm.Model{},
		Name:  "Tanjim Hossain",
		Email: "ths.rupanti@gmail.com",
	})
}

func connectDb() {
	connection := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err == nil {
		database = Database{db}
		database.doMigration()
		database.seedData()
	}
}

func startRouter() error {
	engine := gin.Default()

	userController := user.ProvideUserController(user.ProvideUserService(database.DB))

	userGroup := engine.Group("/users")
	userGroup.GET("", userController.FindAll)
	userGroup.GET("/:id", userController.FindById)

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong!")
	})

	return engine.Run(":8080")
}

func main() {
	connectDb()
	if err := startRouter(); err != nil {
		fmt.Println(err)
	}
}

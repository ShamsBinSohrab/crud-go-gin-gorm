package main

import (
	"crud-api/database"
	"crud-api/router"
	"crud-api/user"
)

func main() {
	db := database.InitDB()
	en := router.InitRouter()

	userAPI := user.InitController(user.InitService(db))
	userAPI.Register(en)

	router.Start(en)
}

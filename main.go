package main

import (
	"crud-api/database"
	"crud-api/router"
	"crud-api/wire"
)

func main() {
	db := database.InitDB()
	en := router.InitRouter()

	userAPI := wire.InitUserAPI(db)
	userAPI.Register(en)

	router.Start(en)
}

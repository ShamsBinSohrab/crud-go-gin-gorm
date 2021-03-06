// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"crud-api/user"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitUserAPI(db *gorm.DB) user.Controller {
	service := user.InitService(db)
	controller := user.InitController(service)
	return controller
}

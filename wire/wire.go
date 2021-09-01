package wire

import (
	"crud-api/user"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) user.Controller {
	wire.Build(user.InitService, user.InitController)
	return user.Controller{}
}

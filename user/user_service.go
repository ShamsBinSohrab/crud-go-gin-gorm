package user

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func InitService(db *gorm.DB) Service {
	return Service{DB: db}
}

func (svc *Service) FindAll() []User {
	var users []User
	svc.DB.Find(&users)
	return users
}

func (svc *Service) FindById(id uint) User {
	var user User
	svc.DB.First(&user, id)
	return user
}

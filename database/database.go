package database

import (
	"crud-api/user"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	username = "aiaibot"
	password = "lena"
	dbname   = "testdb"
	port     = 5432
)

func InitDB() *gorm.DB {
	connection := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, username, password, dbname, port)
	if db, err := gorm.Open(postgres.Open(connection), &gorm.Config{}); err == nil {
		doMigration(db)
		seedData(db)
		return db
	}
	return nil
}

func doMigration(db *gorm.DB) {
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		fmt.Println(err)
	}
}

func seedData(db *gorm.DB) {
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

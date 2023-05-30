package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/risqiikhsani/try-goecho-2/models"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "username:password@tcp(127.0.0.1:3306)/percobaan?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")
	// db.AutoMigrate(&User{})
	// use models
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Account{})

	return db
}

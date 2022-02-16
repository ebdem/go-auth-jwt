package database

import (
	"github.com/ebdem/golang/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// connect to database

	connection, err := gorm.Open(mysql.Open("root:12345678@/auth"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}

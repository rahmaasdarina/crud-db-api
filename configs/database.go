package config

import (
	"crud-api-homework/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func InitDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/hw_crud?charset=utf8mb4&parseTime=True&loc=Local"
	DbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection Opened to Database")
	DbConn.AutoMigrate(&models.Movies{})
	fmt.Println("Database Migrated")
}

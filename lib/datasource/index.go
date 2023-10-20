package datasource

import (
	model "todo/services/todos/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := "root:1234@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic("Failed to connect to database")
	}

	DB = db

	return db
}

func AutoMigrate() {
	err := DB.AutoMigrate(&model.Todo{})
	if err != nil {
		return
	}
}

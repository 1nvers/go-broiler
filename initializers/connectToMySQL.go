package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"fmt"
)

var DB *gorm.DB

func ConnectDatabase(){
	var err error
	LoadEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
  	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!= nil {
		panic("Failed to connect to db")
	}
}
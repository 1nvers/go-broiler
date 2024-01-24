package initializers

import "github.com/1nvers/go-broiler/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}
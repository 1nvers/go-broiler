package database

import "github.com/oneaushaf/go-broiler/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}
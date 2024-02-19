package database

import "github.com/oneaushaf/go-broiler/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Farm{})
	DB.AutoMigrate(&models.Ranch{})
	DB.AutoMigrate(&models.Batch{})
	DB.AutoMigrate(&models.Weighing{})
}

func SetRelation(){
	DB.Model(&models.Farm{}).Association("Ranches")
}
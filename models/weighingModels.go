package models

import (
	"gorm.io/gorm"
)

type Weighing struct {
	gorm.Model
	Code 	  string   `gorm:"unique;not null"`	
	Adress 	  string  	
	Ranches   []Ranch
}
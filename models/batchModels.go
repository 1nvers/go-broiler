package models

import (
	"gorm.io/gorm"
)

type Batch struct {
	gorm.Model
	Count	  uint
	Adress 	  string  	
	Ranches   []Ranch
}
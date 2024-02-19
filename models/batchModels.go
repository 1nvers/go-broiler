package models

import (
	"gorm.io/gorm"
)

type Batch struct {
	gorm.Model
	Count	  uint
	Died	  uint
	Adress 	  string  	
	RanchID	  uint
}
package models

import (
	"gorm.io/gorm"
)

type Batch struct {
	gorm.Model
	InitialQty	  uint
	CurrentQty	  uint
	Finished      bool  	
	RanchID	      uint
}
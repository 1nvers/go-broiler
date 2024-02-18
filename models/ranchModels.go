package models

import (
	"gorm.io/gorm"
)

type Ranch struct {
	gorm.Model
	Capacity  uint  	
	FarmID	  uint 
	Batches	  []Batch
}
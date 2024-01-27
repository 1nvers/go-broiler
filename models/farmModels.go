package models

import (
	"gorm.io/gorm"
)

type Farm struct {
	gorm.Model
	Code string
	LastName  string
	Phone     string
	Email 	  string  `gorm:"unique;not null"`	
}
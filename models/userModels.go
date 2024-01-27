package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Phone     string
	Email 	  string  `gorm:"unique;not null"`
	Password  string
	UserType  string	
}
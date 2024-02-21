package models

import (
	"gorm.io/gorm"
)

type Farm struct {
	gorm.Model
	Code    string `gorm:"unique;not null"`
	Adress  string
	Ranches []Ranch
}

package models

import (
	"gorm.io/gorm"
)

type Ranch struct {
	gorm.Model
	Code     string `gorm:"unique;not null"`
	Capacity uint
	FarmID   uint
	Batches  []Batch
}

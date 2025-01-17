package models

import (
	"gorm.io/gorm"
)

type Weighing struct {
	gorm.Model
	Image         string
	Age           uint
	BatchID       uint
	Deceased      uint
	AverageWeight float64 `gorm:"type:DOUBLE"`
}

package models

import "gorm.io/gorm"

type Expenses struct {
	gorm.Model
	CabangID uint `gorm:"index"`
	Total uint `gorm:"not null"`
}
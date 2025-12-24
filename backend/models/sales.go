package models

import "gorm.io/gorm"

type Sales struct {
	gorm.Model
	CabangID uint `gorm:"index"`
	TotalAmount uint `gorm:"not null"`
	AmountPaid uint	`gorm:"not null"`
	AmountOwed uint `gorm:"not null"`
}
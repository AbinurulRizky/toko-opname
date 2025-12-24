package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string `gorm:"not null"`
	Stock uint `gorm:"not null"`
	Price uint `gorm:"not null"`
}
package models

import "gorm.io/gorm"

type Cabang struct {
	gorm.Model
	Name string `gorm:"not null"`
	Location string `gorm:"type:varchar(255);not null"`
}
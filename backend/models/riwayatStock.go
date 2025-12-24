package models

import (
	"gorm.io/gorm"
)

type RiwayatStock struct {
	gorm.Model
	CabangID uint `gorm:"index"`
	ProductID uint `gorm:"index"`
	UserID uint `gorm:"index"`
	RemainingStock uint `gorm:"not null"`
}
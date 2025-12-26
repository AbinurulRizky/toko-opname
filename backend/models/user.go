package models

import (
	"gorm.io/gorm"
)

const OwnerRole = "Owner"
const EmployeeRole = "Employee"

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Email string `gorm:"type:varchar(255);uniqueIndex;not null"`
	HashedPassword string `gorm:"not null"`
	// 2 role options: 'Owner' and 'Employee'
	Role string `gorm:"type:varchar(50);not null;default:'Employee'"`
	CabangID *uint `gorm:"index"`

	Cabang Cabang `gorm:"foreignKey:CabangID"`
}
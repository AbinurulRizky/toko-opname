package models

import "gorm.io/gorm"

type EmployeeSalary struct {
	gorm.Model
	CabangID uint `gorm:"index"`
	EmployeeName string `gorm:"not null"`
	TotalSalary uint `gorm:"not null"`
}
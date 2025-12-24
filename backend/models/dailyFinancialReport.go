package models

import "gorm.io/gorm"

type DailyFinancialReport struct {
	gorm.Model
	CabangID uint `gorm:"index"`
	ShopMoney uint `gorm:"not null"`
	UnpaidExpenses uint `gorm:"not null"`
	EmployeeSalary uint `gorm:"not null"`
	SalesProceed uint `gorm:"not null"`
	InitialCapital uint `gorm:"not null"`
	ShoppingCapital uint `gorm:"not null"`
	Profit uint `gorm:"not null"`
	NextCapital uint `gorm:"not null"`
}
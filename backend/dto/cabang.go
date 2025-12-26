package dto

type CreateCabang struct {
	Name string `json:"cabang_name" binding:"required"`
	Location string `json:"location" binding:"required"`
}
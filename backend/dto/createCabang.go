package dto

type CreateCabang struct {
	Name string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}
package dto

type CreateCabang struct {
	Name string `json:"cabang_name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

type ShowCabang struct {
	ID uint `json:"id"`
	Name string `json:"name_cabang"`
	Location string `json:"location"`
}
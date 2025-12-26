package dto

type CreateProduct struct {
	ProductName string `json:"product_name"`
	Stock uint `json:"stock"`
	Price uint `json:"price"`
}
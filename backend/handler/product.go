package handler

import (
	"net/http"

	"toko-opname/dto"
	"toko-opname/models"
	"toko-opname/config"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var req dto.CreateProduct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	product := models.Product{
		ProductName: req.ProductName,
		Stock: req.Stock,
		Price: req.Price,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed create the product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success create the product!"})
}
package handler

import (
	"net/http"
	"toko-opname/dto"
	"toko-opname/models"
	"toko-opname/config"

	"github.com/gin-gonic/gin"
)

func CreateCabang(c *gin.Context) {
	var req dto.CreateCabang
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cabang := models.Cabang{
		Name: req.Name,
		Location: req.Location,
	}

	if  err := config.DB.Create(&cabang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cabang."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success create cabang"})
}

package handler

import (
	"toko-opname/config"
	"toko-opname/dto"
	"toko-opname/models"

	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowEmployee(c *gin.Context) {
	users := []models.User{}
	db := config.DB

	result := db.Where("role = ?", "Employee").Preload("Cabang").Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data."})
		return
	}

	var response []dto.ShowEmployee
	for _, u := range users {
		res := dto.ShowEmployee {
			ID:			u.ID,
			Username: 	u.Username,
			Email:		u.Email,
			Role:		u.Role,
			CabangID: 	u.CabangID,
			CreatedAt: 	u.CreatedAt,
		}

		if u.Cabang.ID != 0 {
			res.Cabang = &dto.CabangResponse{
				ID: u.Cabang.ID,
				Name: u.Cabang.Name,
				Location: u.Cabang.Location,
			}
		}

		response = append(response, res)
	}

	c.JSON(http.StatusOK, response)
}
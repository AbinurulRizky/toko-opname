package middleware

import (
	"net/http"

	"toko-opname/models"
	"toko-opname/config"

	"github.com/gin-gonic/gin"
)

func EmployeeOnly(c *gin.Context) {
	UserID, exist := c.Get("UserID")
	if !exist{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No context user not found"})
		c.Abort()
		return
	}

	var user models.User
	if err := config.DB.First(&user, UserID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	if user.Role != models.EmployeeRole {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Employee role is required"})
		c.Abort()
		return
	}

	c.Next()
}
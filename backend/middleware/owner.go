package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"toko-opname/config"
	"toko-opname/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenString, &models.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(*models.JWTCustomClaims); ok && token.Valid {
		c.Set("UserID", claims.UserID)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token", "details": err.Error()})
		return
	}
}

func OwnerOnly(c *gin.Context) {
	UserID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No user context found."})
		c.Abort()
		return
	}

	var user models.User
	if err := config.DB.First(&user, UserID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	if user.Role != models.OwnerRole {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Owner access required."})
		c.Abort()
		return
	}

	c.Next()
}
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"toko-opname/config"

	"os"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		panic("Failed to load .env file")
	}
	
	config.InitDB()

	port := os.Getenv("PORT")
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {"message": "pong"})
	})
	server.Run(":" + port)
}
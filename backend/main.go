package main

import (
	"net/http"
	"toko-opname/config"
	"toko-opname/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

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
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			c.Next()
		}
	})
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {"message": "pong"})
	})
	apiRoutes := server.Group("/api")
	routes.AuthRoutes(apiRoutes)
	routes.OwnerRoutes(apiRoutes)
	server.Run(port)
}
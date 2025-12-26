package routes

import (
	"toko-opname/handler"
	"toko-opname/middleware"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.RouterGroup) {
	router.POST("/employee/create-product", middleware.AuthMiddleware, middleware.EmployeeOnly, handler.CreateProduct)
}
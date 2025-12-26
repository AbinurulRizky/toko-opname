package routes

import (
	"toko-opname/handler"
	"toko-opname/middleware"

	"github.com/gin-gonic/gin"
)

func OwnerRoutes(router *gin.RouterGroup) {
	router.POST("/register", middleware.AuthMiddleware, middleware.OwnerOnly, handler.RegisterUser)
	router.POST("/owner/createcabang", middleware.AuthMiddleware, middleware.OwnerOnly, handler.CreateCabang)
}
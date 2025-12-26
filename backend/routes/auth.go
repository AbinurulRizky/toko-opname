package routes

import (
	"toko-opname/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", handler.LoginUser)
}
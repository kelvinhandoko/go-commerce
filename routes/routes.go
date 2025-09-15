package routes

import (
	"go-commerce/cmd/user/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler handler.UserHandler) {
	router.GET("/ping", userHandler.Ping)
	router.POST("/register", userHandler.Register)
}

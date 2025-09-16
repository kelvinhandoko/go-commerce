package routes

import (
	"go-commerce/cmd/user/handler"
	"go-commerce/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler handler.UserHandler, jwtSecret string) {
	router.Use(middleware.RequestLogger())
	router.GET("/ping", userHandler.Ping)
	router.POST("/v1/register", userHandler.Register)
	router.POST("/v1/login", userHandler.Login)

	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	private := router.Group("/api")
	private.Use(authMiddleware)
	private.GET("/v1/userInfo", userHandler.GetUserInfo)
}

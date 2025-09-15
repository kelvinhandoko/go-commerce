package main

import (
	"go-commerce/cmd/user/handler"
	"go-commerce/cmd/user/repository"
	"go-commerce/cmd/user/resource"
	"go-commerce/cmd/user/service"
	"go-commerce/cmd/user/usecase"
	"go-commerce/config"
	"go-commerce/infrastucture/log"
	"go-commerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	redis := resource.InitRedis(&cfg)
	db := resource.InitDB(&cfg)

	log.SetupLogger()

	userRepository := repository.NewUserRepository(db, redis)
	userService := service.NewUserService(*userRepository)
	userUseCase := usecase.NewUserUsecase(*userService)
	userHandler := handler.NewUserHandler(*userUseCase)

	port := cfg.App.Port
	router := gin.Default()
	routes.SetupRoutes(router, *userHandler)
	router.Run(":" + port)

	log.Logger.Infof("Server running on port %s", port)
}

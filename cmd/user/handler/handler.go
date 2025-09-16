package handler

import (
	"go-commerce/cmd/user/usecase"
	"go-commerce/infrastucture/log"
	"go-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var param models.RegisterParameter
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Logger.Error("invalid parameter", err)
		c.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
	}

	if len(param.Password) < 6 || param.Password != param.ConfirmPassword {
		log.Logger.Error(param.Password)
		log.Logger.Error(param.ConfirmPassword)
		log.Logger.Error("password and confirm password do not match")
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "password and confirm password do not match"})
		return
	}

	user, err := h.UserUsecase.GetUserByEmail(c.Request.Context(), param.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	if user.ID != 0 {
		log.Logger.Error("email already registered")
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "email already registered"})
		return
	}

	if err := h.UserUsecase.CreateUser(&models.User{
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	}); err != nil {
		log.Logger.Error("failed to create user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "register success"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var param models.LoginParameter

	if err := c.ShouldBindJSON(&param); err != nil {
		log.Logger.Error("invalid parameter", err)
		c.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
		return
	}
	token, err := h.UserUsecase.Login(c.Request.Context(), param)

	if err != nil {
		log.Logger.Error("failed to login", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error_message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userIDStr, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error_message": "unauthorized"})
		return
	}

	userID, ok := userIDStr.(float64)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error_message": "unauthorized"})
		return
	}

	user, err := h.UserUsecase.GetUserById(c.Request.Context(), int64(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": user.Name, "email": user.Email})
}

func (h *UserHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

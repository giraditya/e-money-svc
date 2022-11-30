package controllers

import (
	"emoney-service/presentation"
	"emoney-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
}

type authController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		AuthService: authService,
	}
}

func (u *authController) Login(c *gin.Context) {
	var request presentation.AuthLoginRequest
	var response presentation.AuthLoginResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := u.AuthService.Login(c.Request.Context(), request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Token = result.Token
	response.ExpiredAt = result.ExpiredAt

	c.JSON(http.StatusOK, gin.H{"data": response})
}

package controllers

import (
	"emoney-service/helpers"
	"emoney-service/presentation"
	"emoney-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	c.ShouldBindJSON(&request)

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		helpers.ErrorValidate(c, http.StatusBadRequest, err)
		return
	}

	result, err := u.AuthService.Login(c.Request.Context(), request.Username, request.Password)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.Token = result.Token
	response.ExpiredAt = result.ExpiredAt

	helpers.SuccessResponse(c, http.StatusOK, "Login Success", response)
}

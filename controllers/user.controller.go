package controllers

import (
	"emoney-service/helpers"
	"emoney-service/presentation"
	"emoney-service/service"
	"emoney-service/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	Register(c *gin.Context)
	FetchCurrentUser(c *gin.Context)
}

type userController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		UserService: userService,
	}
}

func (u *userController) Register(c *gin.Context) {
	var request presentation.UserRegisterRequest
	var response presentation.UserRegisterResponse
	c.ShouldBindJSON(&request)

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		helpers.ErrorValidate(c, http.StatusBadRequest, err)
		return
	}

	_, err := u.UserService.Register(c.Request.Context(), request.Email, request.Username, request.Password)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.Status = "OK"

	helpers.SuccessResponse(c, http.StatusOK, "Register Success", response)
}

func (u *userController) FetchCurrentUser(c *gin.Context) {
	var response presentation.UserFetchResponse
	id, err := token.ExtractTokenID(c)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	res, err := u.UserService.FetchByUserID(c.Request.Context(), id)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	response.Email = res.Email
	response.Username = res.Username
	response.ID = res.ID

	helpers.SuccessResponse(c, http.StatusOK, "Current User", response)
}

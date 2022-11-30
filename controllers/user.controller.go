package controllers

import (
	"emoney-service/presentation"
	"emoney-service/service"
	"emoney-service/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
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

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := u.UserService.Register(c.Request.Context(), request.Email, request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Status = "OK"

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (u *userController) FetchCurrentUser(c *gin.Context) {
	var response presentation.UserFetchResponse
	id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := u.UserService.FetchByUserID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response.Email = res.Email
	response.Username = res.Username
	response.ID = res.ID

	c.JSON(http.StatusOK, gin.H{"data": response})
}

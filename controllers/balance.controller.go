package controllers

import (
	"emoney-service/helpers"
	"emoney-service/presentation"
	"emoney-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BalanceController interface {
	FetchBalanceByUserID(c *gin.Context)
}

type balanceController struct {
	BalanceService service.BalanceService
}

func NewBalanceController(balanceService service.BalanceService) BalanceController {
	return &balanceController{
		BalanceService: balanceService,
	}
}

func (u *balanceController) FetchBalanceByUserID(c *gin.Context) {
	var request presentation.BalanceFetchByUserIDRequest
	var response presentation.BalanceFetchByUserIDResponse
	c.ShouldBindUri(&request)

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		helpers.ErrorValidate(c, http.StatusBadRequest, err)
		return
	}

	userID, _ := strconv.ParseUint(request.UserID, 10, 32)
	result, err := u.BalanceService.FetchByUserID(c.Request.Context(), uint(userID))
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.UserID = result.UserID
	response.Balance = result.Balance

	helpers.SuccessResponse(c, http.StatusOK, "Balance", response)
}

package controllers

import (
	"emoney-service/presentation"
	"emoney-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := strconv.ParseUint(request.UserID, 10, 32)
	result, err := u.BalanceService.FetchByUserID(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.UserID = result.UserID
	response.Balance = result.Balance

	c.JSON(http.StatusOK, gin.H{"data": response})
}

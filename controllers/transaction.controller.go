package controllers

import (
	"emoney-service/presentation"
	"emoney-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	FetchInquiry(c *gin.Context)
	FetchHistoryByUserID(c *gin.Context)
	Confirm(c *gin.Context)
}

type transactionController struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &transactionController{
		TransactionService: transactionService,
	}
}

func (u *transactionController) FetchInquiry(c *gin.Context) {
	res, err := u.TransactionService.FetchInquiry(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (u *transactionController) FetchHistoryByUserID(c *gin.Context) {
	var request presentation.TransactionFetchHistoryByUserIDRequest
	var response []presentation.TransactionFetchHistoryByUserIDResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := u.TransactionService.FetchHistoryByUserID(c.Request.Context(), request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, v := range res {
		response = append(response, presentation.TransactionFetchHistoryByUserIDResponse{
			UserID:      v.UserID,
			Category:    v.Biller.Category,
			Product:     v.Biller.Product,
			Description: v.Biller.Description,
			Price:       v.Biller.Price,
			Fee:         v.Biller.Fee,
			Status:      v.Status,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (u *transactionController) Confirm(c *gin.Context) {
	var request presentation.TransactionConfirmRequest
	var response presentation.TransactionConfirmResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := u.TransactionService.Confirm(c.Request.Context(), request.UserID, request.BillerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Status = "OK"
	c.JSON(http.StatusOK, gin.H{"data": response})
}

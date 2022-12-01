package controllers

import (
	"emoney-service/helpers"
	"emoney-service/presentation"
	"emoney-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		helpers.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	helpers.SuccessResponse(c, http.StatusOK, "Inquiry Colletion", res)
}

func (u *transactionController) FetchHistoryByUserID(c *gin.Context) {
	var request presentation.TransactionFetchHistoryByUserIDRequest
	var response []presentation.TransactionFetchHistoryByUserIDResponse
	c.ShouldBindUri(&request)

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		helpers.ErrorValidate(c, http.StatusBadRequest, err)
		return
	}

	res, err := u.TransactionService.FetchHistoryByUserID(c.Request.Context(), request.UserID)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err)
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

	helpers.SuccessResponse(c, http.StatusOK, "History User Transaction", response)
}

func (u *transactionController) Confirm(c *gin.Context) {
	var request presentation.TransactionConfirmRequest
	var response presentation.TransactionConfirmResponse
	c.ShouldBindJSON(&request)

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		helpers.ErrorValidate(c, http.StatusBadRequest, err)
		return
	}

	_, err := u.TransactionService.Confirm(c.Request.Context(), request.UserID, request.BillerID)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.Status = "OK"
	helpers.SuccessResponse(c, http.StatusOK, "Confirm Transaction Success", response)
}

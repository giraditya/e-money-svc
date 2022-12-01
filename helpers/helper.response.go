package helpers

import (
	"emoney-service/presentation"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := presentation.SuccessResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	if statusCode >= 400 {
		ctx.AbortWithStatusJSON(statusCode, response)
	} else {
		ctx.JSON(statusCode, response)
	}
}

func ErrorResponse(ctx *gin.Context, statusCode int, Error interface{}) {
	response := presentation.ErrorResponse{
		StatusCode: statusCode,
		Error:      Error,
	}
	ctx.AbortWithStatusJSON(statusCode, response)
}

func ErrorValidate(ctx *gin.Context, statusCode int, err interface{}) {
	response := presentation.ErrorValidateResponse{
		StatusCode: statusCode,
	}
	validationErrors := err.(validator.ValidationErrors)
	for _, v := range validationErrors {
		buildString := fmt.Sprintf("Error: %v With Error: %v", v.Field(), v.Error())
		response.Error = append(response.Error, buildString)
	}
	ctx.AbortWithStatusJSON(statusCode, response)
}

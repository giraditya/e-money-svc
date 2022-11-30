package routes

import (
	"emoney-service/controllers"
	"emoney-service/middlewares"
	"emoney-service/repository"
	"emoney-service/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitTransactionRoutes(db *gorm.DB, route *gin.Engine) {
	transactionRepository := repository.NewTransactionRepository()
	userRepository := repository.NewUserRepository()
	balanceRepository := repository.NewBalanceRepository()
	transactionService := service.NewTransactionService(db, transactionRepository, balanceRepository, userRepository)
	transactionController := controllers.NewTransactionController(transactionService)

	groupRoute := route.Group("/v1")
	groupRoute.Use(middlewares.JWTAuthMiddleware())
	groupRoute.GET("/transaction/inquiry/", transactionController.FetchInquiry)
	groupRoute.GET("/transaction/history/:userid", transactionController.FetchHistoryByUserID)
	groupRoute.POST("/transaction/confirm", transactionController.Confirm)
}

package routes

import (
	"emoney-service/controllers"
	"emoney-service/middlewares"
	"emoney-service/repository"
	"emoney-service/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitBalanceRoutes(db *gorm.DB, route *gin.Engine) {
	balanceRepository := repository.NewBalanceRepository()
	balanceService := service.NewBalanceService(db, balanceRepository)
	balanceController := controllers.NewBalanceController(balanceService)

	groupRoute := route.Group("/v1")
	groupRoute.Use(middlewares.JWTAuthMiddleware())
	groupRoute.GET("/balance/:userid", balanceController.FetchBalanceByUserID)
}

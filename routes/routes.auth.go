package routes

import (
	"emoney-service/controllers"
	"emoney-service/repository"
	"emoney-service/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	userRepository := repository.NewUserRepository()
	authService := service.NewAuthService(db, userRepository)
	authController := controllers.NewAuthController(authService)

	groupRoute := route.Group("/v1")
	groupRoute.POST("/auth/login", authController.Login)
}

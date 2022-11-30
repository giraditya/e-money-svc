package routes

import (
	"emoney-service/controllers"
	"emoney-service/middlewares"
	"emoney-service/repository"
	"emoney-service/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository)
	userController := controllers.NewUserController(userService)

	groupRoute := route.Group("/v1")
	groupRoute.POST("/user/register", userController.Register)

	groupRoute.Use(middlewares.JWTAuthMiddleware())
	groupRoute.GET("/user", userController.FetchCurrentUser)
}

package main

import (
	"emoney-service/app"
	"emoney-service/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := SetupRouter()
	log.Fatal(app.Run(":8080"))
}

func SetupRouter() *gin.Engine {
	db := app.ConnectDatabase()
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	routes.InitAuthRoutes(db, app)
	routes.InitUserRoutes(db, app)
	routes.InitBalanceRoutes(db, app)
	routes.InitTransactionRoutes(db, app)

	return app
}

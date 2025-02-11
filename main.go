// Fix the model for hotel: price and data type
package main

import (
	"flockstay_api/controllers"
	"flockstay_api/initializers"
	"flockstay_api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDb()
	initializers.SyncDb()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
	router.GET("/me", middlewares.ValidateToken, controllers.GetMe)
	router.GET("/hotels", controllers.GetHotelList)
	router.Run()
}

package main

import (
	"flockstay_api/controllers"
	"flockstay_api/initializers"
	"flockstay_api/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDb()
	initializers.SyncDb()
}

func main() {
	router := gin.Default()
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
	router.GET("/me", middlewares.ValidateToken, controllers.GetMe)
	router.Run()
}

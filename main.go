package main

import (
	"flockstay_api/controllers"
	"flockstay_api/initializers"

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
	router.Run()
}

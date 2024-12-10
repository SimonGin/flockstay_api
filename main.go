package main

import (
	"flockstay_api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(ginCtx *gin.Context) {
		ginCtx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}

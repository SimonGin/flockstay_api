package controllers

import (
	"flockstay_api/initializers"
	"flockstay_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHotelList(ginCtx *gin.Context) {
	var hotels []models.Hotel
	var count int64
	initializers.DB.Find(&hotels)
	initializers.DB.Model(&models.Hotel{}).Count(&count)
	ginCtx.JSON(http.StatusOK, gin.H{
		"msg": "Successfully retrieved all the hotel",
		"metadata": gin.H{
			"count": count,
			"data":  hotels,
		},
	})
}

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
	initializers.DB.Preload("City").Find(&hotels)
	initializers.DB.Model(&models.Hotel{}).Count(&count)

	type HotelResData struct {
		ID           uint    `json:"id"`
		Name         string  `json:"name"`
		City         string  `json:"city"`
		Address      string  `json:"address"`
		Description  string  `json:"description"`
		Rating       float64 `json:"rating"`
		Images       string  `json:"images"`
		CheckInTime  string  `json:"check_in_time"`
		CheckOutTime string  `json:"check_out_time"`
	}

	var hotelData []HotelResData
	for _, hotel := range hotels {
		hotelData = append(hotelData, HotelResData{
			ID:           hotel.ID,
			Name:         hotel.Name,
			City:         hotel.City.Name,
			Address:      hotel.Address,
			Description:  hotel.Description,
			Rating:       hotel.Rating,
			Images:       hotel.Images,
			CheckInTime:  hotel.CheckInTime,
			CheckOutTime: hotel.CheckOutTime,
		})
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"msg": "Successfully retrieved all the hotel",
		"metadata": gin.H{
			"count": count,
			"data":  hotelData,
		},
	})
}

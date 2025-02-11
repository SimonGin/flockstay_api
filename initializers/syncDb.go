package initializers

import "flockstay_api/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.City{})
	DB.AutoMigrate(&models.Hotel{})
}

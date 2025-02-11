package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"unique;not null;"`
	Image string `json:"image"`
	// Relations
	Hotels []Hotel `json:"hotels" gorm:"foreignKey:CityID"`
}

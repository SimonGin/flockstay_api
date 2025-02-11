package models

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name         string  `json:"name" gorm:"not null;"`
	CityID       uint    `json:"city_id"`
	Address      string  `json:"address" gorm:"not null;"`
	Description  string  `json:"description"`
	Rating       float64 `json:"rating" gorm:"type:DOUBLE PRECISION"`
	Images       string  `json:"images" gorm:"type:jsonb;default:'[]'"`
	CheckInTime  string  `json:"check_in_time" gorm:"not null;default:13:00"`
	CheckOutTime string  `json:"check_out_time" gorm:"not null;default:11:00"`
	// Relations
	City City `gorm:"foreignKey:CityID;references:ID"`
}

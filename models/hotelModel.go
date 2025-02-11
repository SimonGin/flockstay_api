package models

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null;"`
	City         string `json:"city" gorm:"not null;"`
	Address      string `json:"address" gorm:"not null;"`
	Description  string `json:"description"`
	Rating       string `json:"rating"`
	Images       string `json:"images" gorm:"type:jsonb;default:'[]'"`
	CheckInTime  string `json:"check_in_time" gorm:"not null;default:13:00"`
	CheckOutTime string `json:"check_out_time" gorm:"not null;default:11:00"`
}

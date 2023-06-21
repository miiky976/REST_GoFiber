package models

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"desc"`
	Status      bool   `gorm:"default: false" json:"stat"`
}

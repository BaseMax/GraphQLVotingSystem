package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Title string `json:"title"`
	Parint_id uint `json:"parint_id"`
}
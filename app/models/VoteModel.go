package models

import (
	"time"

	"gorm.io/gorm"
)

type Vote struct{
	gorm.Model
	Title       string
	StartDate    time.Time
	EndDate      time.Time
	Questions   Question
	Applications Application
  }
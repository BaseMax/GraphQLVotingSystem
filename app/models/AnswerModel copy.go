package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	questionId  uint32
	applicationId uint32
	Question    Question  
	Application_id uint32
	Answer_id      uint32
}

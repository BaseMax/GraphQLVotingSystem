package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	questionId  uint32
	applicationId uint32
	Question    Question  
	Application Application
	Answer      String
}

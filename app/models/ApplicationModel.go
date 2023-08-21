package models

import (
	"gorm.io/gorm"
)

type Applications struct {
	gorm.Model
	Vote_id uint32
	User_id uint32
	Vote    Vote
	User    User
	Answers Answer
}

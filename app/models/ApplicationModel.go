package models

import (
	"gorm.io/gorm"
)

type Application struct{
	gorm.Model
	VoteId      uint32
	UserId      uint32
	Vote        Vote     
	User        User     
	Answers     Answer
  }
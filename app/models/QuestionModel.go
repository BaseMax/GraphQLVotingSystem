package models

import (

	"gorm.io/gorm"
)

type Question struct{
	gorm.Model
	Text    string
	VoteId  uint32
	Vote    Vote
  }
package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	Gender      UserGender `gorm:"type:enum('man', 'feman','other')"`
	Email       string     `json:"email"`
	VerfiyEmail uint8      `json:"verify-email"`
	CodeVerify  uint8      `json:"code-verify"`
	Role        UserRole   `gorm:"type:enum('admin', 'moderator','supervisor', 'user','question-maker' , 'corrector')"`
	LastIp      string     `json:"last-ip"`
	Name        string     `json:"name"`
	LastName    string     `json:"L-name"`
	Codemile    string     `json:"code-mile"`
}

type UserRole string
type UserGender string

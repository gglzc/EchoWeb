package model

import (
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Password	string  `gorm:"not null"`
	Email		string	`gorm:"not null"`
	Birthday	string	`gorm:"not null"`
	Ban			bool	
}



package model

import (
	"gorm.io/gorm"
)

type Post struct{
	gorm.Model
	Title		string	`gorm:"not null"`
	Article		string	`gorm:"not null"`
	Location	string	`gorm:"not null"`
	Tag			string	`gorm:"not null"`
}
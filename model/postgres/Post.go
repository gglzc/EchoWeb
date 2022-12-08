package model

import (
	"gorm.io/gorm"
)

type Post struct{
	gorm.Model
	Title		string	`gorm:"not null"`
	Article		string	`gorm:"not null"`
	Location	string	`gorm:"not null"`
	Post_Like	uint64	
	UserID		uint64	`gorm:"references:UserID"`	      
}


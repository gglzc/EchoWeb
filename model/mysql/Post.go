package model

import (
	"gorm.io/gorm"
)

type Post struct{
	gorm.Model
	Title		string	`gorm:"not null"`
	Article		string	`gorm:"not null"`
	Location	string	`gorm:"not null"`
<<<<<<< HEAD:model/postgres/Post.go
	Post_Like	uint64	
	UserID		uint64	`gorm:"references:UserID"`	      
=======
	Tag			string  `gorm:"not null"`
	UserID		uint	
>>>>>>> 5971af5 (improve whole structure):model/mysql/Post.go
}


package model

import (

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Username	string  `json:"username" gorm:"not null"`
	Password	string  `json:"password" gorm:"not null"`
	Email		string	`json:"email"    gorm:"not null"`
	Birthday	string	`json:"birthday" gorm:"not null"`
	Status		bool	`json:"status"   gorm:"not null"`	
	Posts		[]Post	`json:"posts"`
}

func (u *User) CheckUser(username string){
	db.Model
}

func (u *User) CheckEmailExist(email string){
	
}


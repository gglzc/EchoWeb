package model

import (
	"time"
)

type User struct{
	UserID        		uint64    `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Username  			string    `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  			string    `gorm:"type:varchar(255) NOT NULL;" json:"password,omitempty"`
	Email	  			string	  `gorm:"type:varchar(255) NOT NULL;" json:"email,omitempty"`
	Status    			int32     `gorm:"type:int(5);" json:"status,omitempty"`
	ActivationToken 	string    `gorm:"varchar(255)"`
    Activated 			int       `gorm:"type:int(10)"`
	CreatedAt 			time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt 			time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}







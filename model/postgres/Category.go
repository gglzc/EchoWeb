package model

import "gorm.io/gorm"

type Category struct{
	gorm.Model
	Name		string `json:"name"` 
	Posts 		[]Post `"gorm":many2many ,Post_category`
} 
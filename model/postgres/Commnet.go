package model

import "gorm.io/gorm"

//留言
type Comment struct{
	gorm.Model
	CommentID       uint64	`gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"comment_id,omitempty"`
	Text 			string  `gorm:"type:bigchar(255) NOT NULL;" json:"text,omitempty"`
	Like			uint64  `gorm:"type:bigint(20) NOT NULL;" json:"like,omitempty"`
	UserID			uint64	 
}
package models

import "github.com/jinzhu/gorm"

// _ "github.com/jinzhu/gorm/dialects/mysql"

// Task model representation
type Task struct {
	gorm.Model
	Name    string `json:"name" xml:"name" form:"name" query:"name"`
	IsReady bool   `gorm:"TINYINT" json:"isReady" xml:"isReady" form:"isReady" query:"isReady"`
	User    User
	UserID  uint
}

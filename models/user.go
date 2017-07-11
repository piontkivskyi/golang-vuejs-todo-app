package models

import "github.com/jinzhu/gorm"

// User model implementation
type User struct {
	gorm.Model
	Name     string
	Username string
	Email    string
	Type     int
	Tasks    []Task `gorm:"ForeignKey:UserID"`
	Password string
}

const (
	_ = iota
	active
	inactive
)

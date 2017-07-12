package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model implementation
type User struct {
	gorm.Model
	Name     string
	Username string
	Tasks    []Task `gorm:"ForeignKey:UserID"`
	Password string
}

// BeforeCreate callback to store password in proper way
func (u *User) BeforeCreate() (err error) {
	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = string(hash)
	return
}

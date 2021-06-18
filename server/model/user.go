package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string
	Email        string
}

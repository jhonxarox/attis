package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Type    string
	Balance float64
	UserID  uint
	User    User
	History []Transaction
}

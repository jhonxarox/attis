package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount    float64
	Timestamp time.Time `gorm:"autoCreateTime"`
	Type      string
	AccountID uint
	Account   Account
}

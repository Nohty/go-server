package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model `json:"-"`
	Street   string `gorm:"not null" json:"street"`
	City     string `gorm:"not null" json:"city"`
	Postcode string `gorm:"not null" json:"postcode"`
	Number   string `gorm:"not null" json:"number"`
}

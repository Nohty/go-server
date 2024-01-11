package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string  `gorm:"uniqueIndex;not null"`
	Email      string  `gorm:"uniqueIndex;not null"`
	Phone      string  `gorm:"not null"`
	Password   string  `gorm:"not null"`
	Permission int64   `gorm:"not null"`
	WalletAddr string  `gorm:"not null"`
	Address    Address `gorm:"foreignKey:ID"`
	Contacts   []User  `gorm:"many2many:user_contacts;"`
}

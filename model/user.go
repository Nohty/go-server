package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Username   string  `gorm:"uniqueIndex;not null" json:"username"`
	Email      string  `gorm:"uniqueIndex;not null" json:"email"`
	Phone      string  `gorm:"not null" json:"phone"`
	Password   string  `gorm:"not null" json:"-"`
	Permission int64   `gorm:"not null" json:"permission"`
	WalletAddr string  `gorm:"not null" json:"wallet_addr"`
	Address    Address `gorm:"foreignKey:ID" json:"address"`
	Contacts   []User  `gorm:"many2many:user_contacts;" json:"contacts"`
}

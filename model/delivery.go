package model

import "gorm.io/gorm"

type Status int64
type Size int64

const (
	Waiting Status = iota
	Paid
	Accepted
	Rejected
	Completed

	Small Size = iota
	Medium
	Large
)

type Delivery struct {
	gorm.Model
	FromAddress Address `gorm:"foreignKey:ID"`
	ToAddress   Address `gorm:"foreignKey:ID"`
	FromUser    User    `gorm:"foreignKey:ID"`
	ToUser      User    `gorm:"foreignKey:ID"`
	Driver      User    `gorm:"foreignKey:ID"`
	Price       float64 `gorm:"not null"`
	Status      Status  `gorm:"not null"`
	Size        Size    `gorm:"not null"`
	Description string  `gorm:"not null"`
}

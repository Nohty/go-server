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
	gorm.Model `json:"-"`
	FromAddress Address `gorm:"foreignKey:ID" json:"from_address"`
	ToAddress   Address `gorm:"foreignKey:ID" json:"to_address"`
	FromUser    User    `gorm:"foreignKey:ID" json:"from_user"`
	ToUser      User    `gorm:"foreignKey:ID" json:"to_user"`
	Driver      User    `gorm:"foreignKey:ID" json:"driver"`
	Price       float64 `gorm:"not null" json:"price"`
	Status      Status  `gorm:"not null" json:"status"`
	Size        Size    `gorm:"not null" json:"size"`
	Description string  `gorm:"not null" json:"description"`
}

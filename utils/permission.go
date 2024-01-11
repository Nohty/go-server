package utils

import (
	"github.com/Nohty/api/database"
	"github.com/Nohty/api/model"
)

const (
	IsAdmin int64 = 1 << iota
	IsDriver
	IsUser
)

func HasPermission(permission int64, flags ...int64) bool {
	for _, flag := range flags {
		if permission&flag == flag {
			return true
		}
	}
	return false
}

func NewPermissionFlags(flags ...int64) int64 {
	var permission int64
	for _, flag := range flags {
		permission = permission | flag
	}
	return permission
}

func AddPermission(permission int64, flags ...int64) int64 {
	for _, flag := range flags {
		permission = permission | flag
	}
	return permission
}

func RemovePermission(permission int64, flags ...int64) int64 {
	for _, flag := range flags {
		permission = permission &^ flag
	}
	return permission
}

func GetPermissionFromDB(userId uint) int64 {
	db := database.DB

	var user model.User
	if err := db.Select("permission").First(&user, userId).Error; err != nil {
		return 0
	}

	return user.Permission
}

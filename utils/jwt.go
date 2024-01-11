package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims string

const (
	Username Claims = "username"
	UserID   Claims = "user_id"
	Expire   Claims = "exp"
)

func getClaims(c *fiber.Ctx, key Claims) any {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims[string(key)]
}

func GetUserID(c *fiber.Ctx) (uint, error) {
	if id, ok := getClaims(c, UserID).(float64); ok {
		return uint(id), nil
	}

	return 0, errors.New("invalid user id")
}

func GetUsername(c *fiber.Ctx) (string, error) {
	if username, ok := getClaims(c, Username).(string); ok {
		return username, nil
	}

	return "", errors.New("invalid username")
}

func GetExpire(c *fiber.Ctx) (int64, error) {
	if expire, ok := getClaims(c, Expire).(float64); ok {
		return int64(expire), nil
	}

	return 0, errors.New("invalid expire")
}

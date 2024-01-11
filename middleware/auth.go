package middleware

import (
	"github.com/Nohty/api/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.JWT_SECRET)},
		ErrorHandler: jwtError,
	})
}

func jwtError(_ *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing or malformed JWT")
	}

	return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired JWT")
}

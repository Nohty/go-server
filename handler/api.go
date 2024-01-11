package handler

import (
	"github.com/Nohty/api/utils"

	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx) error {
	return utils.Response(c, fiber.StatusOK, "success", nil)
}

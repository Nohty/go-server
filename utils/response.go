package utils

import "github.com/gofiber/fiber/v2"

type HttpResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	status := "success"
	if statusCode >= 400 {
		status = "error"
	}

	return c.Status(statusCode).JSON(&HttpResp{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

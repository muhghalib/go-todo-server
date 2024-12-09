package utils

import "github.com/gofiber/fiber/v2"

func CreateError(code int, message ...string) *fiber.Error {
	return fiber.NewError(code, message...)
}

package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ValidateHandler(c *fiber.Ctx) error {
	return c.SendString(strconv.FormatBool(true))
}

package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/token"
	"strconv"
)

func ValidateHandler(c *fiber.Ctx) error {
	return c.SendString(strconv.FormatBool(token.Validate()))
}

package handler

import "github.com/gofiber/fiber/v2"

func GroupInfo(ctx *fiber.Ctx) error {

	return ctx.SendString("Works")
}

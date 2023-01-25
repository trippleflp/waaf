package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type listEntitledGroupBody struct {
	userId string `json:"userId"`
}

func ListEntitledGroups(c *fiber.Ctx) error {

	body := listEntitledGroupBody{}
	err := c.BodyParser(&body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}

	return c.SendString(body.userId)
}

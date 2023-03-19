package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/postgres"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/token"
)

type CreateTokenBody = models.UserIdWrapper[any]

func CreateToken(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(CreateTokenBody)
	err := c.BodyParser(body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}
	client := postgres.GetConnection()
	groups, err := client.GetEntitledFunctionGroups(body.UserId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Could not read entitled function groups")
		return fiber.NewError(fiber.StatusInternalServerError, "Could not read entitled function groups")
	}

	var hashes []string
	for _, groupId := range groups {
		hash, err := client.GetFunctionGroupHash(groupId, c.UserContext())
		if err != nil {
			log.Debug().Err(err).Str("body", string(c.Body())).Msg("Could not create hash of function group")
			return fiber.NewError(fiber.StatusInternalServerError, "Could not create hash of function group")
		}
		hashes = append(hashes, *hash)
	}

	t, _, err := token.CreateToken(body.UserId, hashes)
	if err != nil {
		log.Debug().Err(err).Msg("Could not create JWT of function groups")
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create JWT of function groups")

	}
	log.Debug().Str("token", t).Msg("Token")

	return c.SendString(t)
}

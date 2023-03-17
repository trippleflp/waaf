package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/postgres"
)

type AddFunctionBody struct {
	FunctionGroupName string `json:"functionGroupName"`
	FunctionTag       string `json:"functionTag"`
}

func AddFunction(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(AddFunctionBody)
	err := c.BodyParser(body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}

	client := postgres.GetConnection()

	id, err := client.GetFunctionGroupId(body.FunctionGroupName, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Database query failed")
		return fiber.NewError(fiber.StatusInternalServerError, "Database query failed")
	}
	if id == "" {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("FunctionGroup does not exist")
		return fiber.NewError(fiber.StatusBadRequest, "FunctionGroup does not exist")
	}
	err = client.AddFunction(body.FunctionTag, id, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Function adding failed")
		return fiber.NewError(fiber.StatusInternalServerError, "Function adding failed")
	}

	return nil
}

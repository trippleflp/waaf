package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/authentication/internal/postgres"
)

type loginBody struct {
	Username string `json:"username",omitempty`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}

func LoginHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := loginBody{}
	err := c.BodyParser(&body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}

	con := postgres.GetConnection()
	exists, err := con.Exists(body.Username, body.Email, c.Context())
	return c.SendString("Login")
}

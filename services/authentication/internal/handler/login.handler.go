package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/authentication/internal/postgres"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/authentication/internal/token"
)

type loginBody struct {
	Username string `json:"username",omitempty`
	Password string `json:"password"`
	Email    string `json:"email",omitempty`
}

func LoginHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := loginBody{}
	err := c.BodyParser(&body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}

	if body.Email == "" && body.Username == "" {
		log.Debug().Msg("Neither email nor userName were provided")
		return fiber.NewError(fiber.StatusBadRequest, "Neither email nor userName were provided")
	}

	con := postgres.GetConnection()
	userId, err := con.Authenticate(body.Username, body.Email, hash(body.Password), c.Context())
	if err != nil {
		log.Debug().Err(err).Str("userName", body.Username).Str("email", body.Email).Msg("Wasn't able to authenticate")
		return fiber.ErrUnauthorized
	}

	token, _, err := token.CreateToken(userId)
	if err != nil {
		log.Debug().Err(err).Msg("Could not create token")
		return fiber.ErrInternalServerError
	}

	return c.SendString(token)
}

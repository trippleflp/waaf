package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/authentication/internal/postgres"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/authentication/internal/token"
)

type registerBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RegisterHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := registerBody{}
	err := c.BodyParser(&body)

	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}

	con := postgres.GetConnection()
	exists, err := con.Exists(body.Username, body.Email, c.Context())
	if exists || err != nil {
		log.Debug().Err(err).Str("userName", body.Username).Str("email", body.Email).Msg("Username or email is already taken, or error")
		if err != nil {
			return fiber.ErrInternalServerError
		}
		return fiber.NewError(fiber.StatusBadRequest, "User already exists")
	}

	userId := uuid.NewString()
	log.Debug().Msg("create token")
	token, _, err := token.CreateToken(userId)
	if err != nil {
		log.Debug().Err(err).Msg("Could not create token")
		return fiber.ErrInternalServerError
	}

	err = con.CreateUser(postgres.User{
		Id:       userId,
		UserName: body.Username,
		Email:    body.Email,
		Password: hash(body.Password),
	}, c.Context())
	if err != nil {
		log.Debug().Err(err).Msg("Could not write to database")
		return fiber.ErrInternalServerError
	}

	return c.SendString(token)
}

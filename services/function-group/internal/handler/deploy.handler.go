package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models"
	deployer "gitlab.informatik.hs-augsburg.de/flomon/waaf/services/deployer/model"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/postgres"
	"os"
)

type DeployBody = models.UserIdWrapper[any]

func Deploy(c *fiber.Ctx) error {
	functionGroupId := c.Params("id")

	body := new(DeployBody)
	err := c.BodyParser(body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}

	client := postgres.GetConnection()
	isDeveloper, err := client.IsAtLeastDeveloper(body.UserId, functionGroupId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Database query failed")
		return fiber.NewError(fiber.StatusInternalServerError, "Database query failed")
	}
	if !isDeveloper {
		log.Debug().Str("body", string(c.Body())).Msg("Only admins and devs can add functionGroups")
		return fiber.NewError(fiber.StatusUnauthorized, "Only admins and devs can add functionGroups")
	}

	group, err := client.GetFunctionGroup(functionGroupId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Database query failed")
		return fiber.NewError(fiber.StatusInternalServerError, "Database query failed")
	}

	b := &deployer.DeployHandlerBody{
		Functions: lo.Map[*postgres.Function, string](group.Functions, func(item *postgres.Function, index int) string {
			return item.FunctionTag
		}),
		FunctionGroupName: group.Name,
	}

	url, exists := os.LookupEnv("DEPLOYER_URL")
	if !exists {
		log.Debug().Msg("Deployer url is not set")
		return fiber.NewError(fiber.StatusBadRequest, "Deployer url is not set")
	}
	res, err := req.R().SetBody(b).Post(url)
	if err != nil {
		log.Debug().Msg("Deployment failed")
		return fiber.NewError(fiber.StatusBadRequest, "Deployment failed")
	}
	if res.StatusCode != 200 {
		log.Debug().Msg("Deployment failed")
		return fiber.NewError(fiber.StatusBadRequest, "Deployment failed")
	}

	return c.SendString("Deployment completed")
}

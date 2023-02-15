package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/postgres"

	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models"
)

type GroupInfoBody = models.UserIdWrapper[any]

func GroupInfo(c *fiber.Ctx) error {
	functionGroupId := c.Params("id")

	body := new(GroupInfoBody)
	err := c.BodyParser(body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}
	client := postgres.GetConnection()
	ok, err := client.IsAtLeastReader(body.UserId, functionGroupId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Database query failed")
		return fiber.NewError(fiber.StatusInternalServerError, "Database query failed")
	}
	if !ok {
		log.Debug().Str("body", string(c.Body())).Msg("Only user with at least read permissions can receive functionGroup details")
		return fiber.NewError(fiber.StatusUnauthorized, "Only user with at least read permissions can receive functionGroup details")
	}

	functionGroup, err := client.GetFunctionGroup(functionGroupId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Database query failed")
		return fiber.NewError(fiber.StatusInternalServerError, "Database query failed")
	}
	var userIds []*string
	for _, user := range functionGroup.Users {
		userIds = append(userIds, &user.UserId)
	}

	data := model.FunctionGroup{
		Name:                  functionGroup.Name,
		ID:                    functionGroup.Id,
		UserIds:               userIds,
		AllowedFunctionGroups: nil,
	}
	respData, err := json.Marshal(data)
	if err != nil {
		log.Debug().Err(err).Interface("functionGroup", functionGroup).Msg("Could not marshall functionGroup")
		return fiber.NewError(fiber.StatusInternalServerError, "Could not marshall functionGroup")
	}

	return c.Send(respData)
}

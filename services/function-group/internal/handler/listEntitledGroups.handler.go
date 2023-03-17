package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/postgres"
)

type ListEntitledGroupBody = models.UserIdWrapper[any]

func ListEntitledGroups(c *fiber.Ctx) error {

	body := new(ListEntitledGroupBody)
	err := c.BodyParser(body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}

	client := postgres.GetConnection()
	exist, err := client.CreateUserIfNotExist(body.UserId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Could not create user")
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create user")
	}
	if !exist {
		functionGroups := []any{}
		responseData, err := json.Marshal(functionGroups)
		if err != nil {
			log.Debug().Err(err).Interface("functionGroup", functionGroups).Msg("Could not marshall functionGroups")
			return fiber.NewError(fiber.StatusInternalServerError, "Could not marshall functionGroups")
		}

		return c.Send(responseData)
	}

	functionGroupIds, err := client.GetEntitledFunctionGroups(body.UserId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Could not read entitled function groups")
		return fiber.NewError(fiber.StatusInternalServerError, "Could not read entitled function groups")
	}
	var functionGroups []*model.FunctionGroup

	for _, groupId := range functionGroupIds {
		groupRaw, err := client.GetFunctionGroup(groupId, c.UserContext())
		if err != nil {
			log.Debug().Err(err).Str("body", string(c.Body())).Msg("Could not build function groupId model")
			return fiber.NewError(fiber.StatusInternalServerError, "Could not build function groupId model")

		}
		functionGroups = append(functionGroups, &model.FunctionGroup{
			Name: groupRaw.Name,
			ID:   groupRaw.Id,
			UserIds: lo.Map[*postgres.FunctionGroupToUserRolePair, *string](groupRaw.Users, func(user *postgres.FunctionGroupToUserRolePair, index int) *string {
				return &user.UserId
			}),
		})
	}

	responseData, err := json.Marshal(functionGroups)
	if err != nil {
		log.Debug().Err(err).Interface("functionGroup", functionGroups).Msg("Could not marshall functionGroups")
		return fiber.NewError(fiber.StatusInternalServerError, "Could not marshall functionGroups")
	}

	return c.Send(responseData)
}

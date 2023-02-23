package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/postgres"

	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models"
)

type EditUserRoleBody = models.UserIdWrapper[*model.UserRolePairInput]

func EditUserRole(c *fiber.Ctx) error {
	c.Accepts("application/json")
	functionGroupId := c.Params("id")

	body := new(EditUserRoleBody)
	err := c.BodyParser(body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}

	client := postgres.GetConnection()
	exists, err := client.FunctionGroupExists(functionGroupId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Database query failed")
		return fiber.NewError(fiber.StatusInternalServerError, "Database query failed")
	}
	if !exists {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("FunctionGroup does not exist")
		return fiber.NewError(fiber.StatusBadRequest, "FunctionGroup does not exist")
	}

	isAdmin, err := client.IsAdmin(body.UserId, functionGroupId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Database query failed")
		return fiber.NewError(fiber.StatusInternalServerError, "Database query failed")
	}
	if !isAdmin {
		log.Debug().Str("body", string(c.Body())).Msg("Only admins can edit user roles")
		return fiber.NewError(fiber.StatusUnauthorized, "Only admins can edit user roels")
	}

	functionGroup, err := client.EditUserRole(body.Data, functionGroupId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("User role editing failed")
		return fiber.NewError(fiber.StatusInternalServerError, "User role editing failed")
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

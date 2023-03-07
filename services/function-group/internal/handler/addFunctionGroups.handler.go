package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/postgres"
)

type AddFunctionGroupsBody = models.UserIdWrapper[[]string]

func AddFunctionGroups(c *fiber.Ctx) error {

	c.Accepts("application/json")
	functionGroupId := c.Params("id")

	body := new(AddFunctionGroupsBody)
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

	isAdmin, err := client.IsAtLeastDeveloper(body.UserId, functionGroupId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Database query failed")
		return fiber.NewError(fiber.StatusInternalServerError, "Database query failed")
	}
	if !isAdmin {
		log.Debug().Str("body", string(c.Body())).Msg("Only admins and devs can add functionGroups")
		return fiber.NewError(fiber.StatusUnauthorized, "Only admins and devs can add functionGroups")
	}

	newlyAddedGroup, alreadyAddedGroup, notExistGroup, err := client.AddFunctionGroups(body.Data, functionGroupId, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("User adding failed")
		return fiber.NewError(fiber.StatusInternalServerError, "User adding failed")
	}

	log.Debug().Err(err).Interface("newlyAddedUser", newlyAddedGroup).Interface("alreadyAddedUsers", alreadyAddedGroup).Interface("notExistGroup", notExistGroup).Msg("Groups added")

	functionGroup, err := client.GetFunctionGroup(functionGroupId, c.UserContext())
	var userIds []*string
	for _, user := range functionGroup.Users {
		userIds = append(userIds, &user.UserId)
	}
	var allowedFunctionGroupIds []*model.FunctionGroupID
	for _, group := range functionGroup.AllowedFunctionGroups {
		allowedFunctionGroupIds = append(allowedFunctionGroupIds, &model.FunctionGroupID{
			Name: &group.ChildFunctionGroup.Name,
			ID:   &group.ChildFunctionGroup.Id,
		})
	}

	data := model.FunctionGroup{
		Name:                  functionGroup.Name,
		ID:                    functionGroup.Id,
		UserIds:               userIds,
		AllowedFunctionGroups: allowedFunctionGroupIds,
	}
	respData, err := json.Marshal(data)
	if err != nil {
		log.Debug().Err(err).Interface("functionGroup", functionGroup).Msg("Could not marshall functionGroup")
		return fiber.NewError(fiber.StatusInternalServerError, "Could not marshall functionGroup")
	}

	return c.Send(respData)
}

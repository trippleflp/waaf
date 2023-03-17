package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/postgres"
)

type CreateGroupBody = models.UserIdWrapper[model.CreateFunctionGroupInput]

func CreateGroup(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(CreateGroupBody)
	err := c.BodyParser(body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}
	allowedFunctionGroups := body.Data.AllowedFunctionGroups
	for _, group := range allowedFunctionGroups {
		if group.ID == nil && group.Name == nil {
			log.Debug().Str("body", string(c.Body())).Msg("Neither the groupId nor the groupName was provided")
			return fiber.NewError(fiber.StatusBadRequest, "Neither the groupId nor the groupName was provided")
		}
	}

	client := postgres.GetConnection()
	groupId, err := client.CreateFunctionGroup(body.UserId, body.Data.GroupName, c.UserContext())
	if err != nil {
		log.Debug().Err(err).Msg("Was not able to write to db")
		return fiber.NewError(fiber.StatusBadRequest, "Was not able to write to db")
	}

	responseData := model.FunctionGroup{
		Name: body.Data.GroupName,
		ID:   *groupId,
		UserIds: []*string{
			&body.UserId,
		},
		//AllowedFunctionGroups: []*model.FunctionGroupID{
		//	{
		//		Name: body.Data.AllowedFunctionGroups[0].Name,
		//		ID:   body.Data.AllowedFunctionGroups[0].ID,
		//	},
		//},
	}

	data, err := json.Marshal(responseData)
	if err != nil {
		log.Debug().Err(err).Interface("functionGroup", responseData).Msg("Could not marshall functionGroup")
		return fiber.NewError(fiber.StatusInternalServerError, "Could not marshall functionGroup")
	}

	return c.Send(data)
}

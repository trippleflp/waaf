package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
)

func (c *pgConnection) CreateFunctionGroup(userId, groupName string, ctx context.Context) (*string, error) {
	groupId := uuid.NewString()
	if _, err := c.db.NewInsert().Model(&User{Id: userId}).Exec(ctx); err != nil {
		return nil, err
	}
	if _, err := c.db.NewInsert().
		Model(&FunctionGroup{
			Id:   groupId,
			Name: groupName,
		}).
		Exec(ctx); err != nil {
		return nil, err
	}
	if _, err := c.db.NewInsert().
		Model(
			&FunctionGroupToUserRolePair{
				UserId:          userId,
				Role:            model.UserRoleAdmin,
				FunctionGroupId: groupId}).
		Exec(ctx); err != nil {
		return nil, err
	}

	return &groupId, nil
}

func (c *pgConnection) AddUsers(users []*model.AddUserToFunctionGroupInput, groupId string, ctx context.Context) ([]string, []string, error) {
	var alreadyAdded []string
	var newlyAdded []string

	for _, user := range users {
		functionGroup := new(FunctionGroup)
		err := c.db.NewSelect().Model(functionGroup).Relation("Users", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("user_id = uuid(?)", user.UserID)
		}).Scan(ctx)
		if err != nil {
			return nil, nil, err
		}
		if len(functionGroup.Users) >= 1 {
			alreadyAdded = append(alreadyAdded, user.UserID)
			continue
		}

		if _, err := c.db.NewInsert().Model(&User{Id: user.UserID}).Ignore().Exec(ctx); err != nil {
			return nil, nil, err
		}
		if _, err := c.db.NewInsert().
			Model(
				&FunctionGroupToUserRolePair{
					UserId:          user.UserID,
					Role:            user.Role,
					FunctionGroupId: groupId}).
			Exec(ctx); err != nil {
			return nil, nil, err
		}
		newlyAdded = append(newlyAdded, user.UserID)
	}
	return newlyAdded, alreadyAdded, nil

}

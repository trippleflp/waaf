package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
)

func (c *PgConnection) CreateFunctionGroup(userId, groupName string, ctx context.Context) (*string, error) {
	groupId := uuid.NewString()
	if _, err := c.CreateUserIfNotExist(userId, ctx); err != nil {
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

func (c *PgConnection) AddUsers(users []*model.UserRolePairInput, groupId string, ctx context.Context) ([]string, []string, error) {
	var alreadyAdded []string
	var newlyAdded []string

	for _, user := range users {
		functionGroup := new(FunctionGroup)
		err := c.db.NewSelect().
			Model(functionGroup).
			Where("id = uuid(?)", groupId).
			Relation("Users", func(q *bun.SelectQuery) *bun.SelectQuery {
				return q.Where("user_id = uuid(?)", user.UserID)
			}).Scan(ctx)
		if err != nil {
			return nil, nil, err
		}
		if len(functionGroup.Users) >= 1 {
			alreadyAdded = append(alreadyAdded, user.UserID)
			continue
		}

		if _, err := c.CreateUserIfNotExist(user.UserID, ctx); err != nil {
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

func (c *PgConnection) AddFunctionGroups(functionGroupIds []string, targetFunctionGroup string, ctx context.Context) ([]string, []string, []string, error) {
	var alreadyAdded []string
	var newlyAdded []string
	var notExist []string

	for _, fnGroupId := range functionGroupIds {
		functionGroup := new(FunctionGroup)

		exists, err := c.db.NewSelect().
			Where("id = uuid(?)", targetFunctionGroup).
			Exists(ctx)

		if err != nil {
			return nil, nil, nil, err
		}
		if !exists {
			notExist = append(notExist, fnGroupId)
			continue
		}

		err = c.db.NewSelect().
			Model(functionGroup).
			Where("id = uuid(?)", targetFunctionGroup).
			Relation("AllowedFunctionGroupIds", func(q *bun.SelectQuery) *bun.SelectQuery {
				return q.Where("parent_function_group_id = uuid(?)", fnGroupId)
			}).Scan(ctx)
		if err != nil {
			return nil, nil, nil, err
		}
		if len(functionGroup.Users) >= 1 {
			alreadyAdded = append(alreadyAdded, fnGroupId)
			continue
		}

		if _, err := c.db.NewInsert().
			Model(
				&AllowedFunctionGroupPair{
					ParentFunctionGroupId: targetFunctionGroup,
					ChildFunctionGroupId:  fnGroupId,
				}).
			Exec(ctx); err != nil {
			return nil, nil, nil, err
		}
		newlyAdded = append(newlyAdded, fnGroupId)
	}
	return newlyAdded, alreadyAdded, notExist, nil

}

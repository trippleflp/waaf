package postgres

import (
	"context"
	"github.com/google/uuid"
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
				Role:            AdminRole,
				FunctionGroupId: groupId}).
		Exec(ctx); err != nil {
		return nil, nil
	}

	return &groupId, nil
}

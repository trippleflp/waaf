package postgres

import (
	"context"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

func (c *PgConnection) IsAdmin(userId, groupId string, ctx context.Context) (bool, error) {
	exists, err := c.db.NewSelect().
		Model((*User)(nil)).
		Where("id = uuid(?)", userId).
		Relation("FunctionGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("function_group_id = uuid(?)", groupId)
		}).
		Exists(ctx)
	return exists, err
}

func (c *PgConnection) FunctionGroupExists(groupId string, ctx context.Context) (bool, error) {
	exists, err := c.db.NewSelect().
		Model((*FunctionGroup)(nil)).
		Where("id = uuid(?)", groupId).
		Exists(ctx)
	return exists, err
}

func (c *PgConnection) GetFunctionGroup(groupId string, ctx context.Context) (*FunctionGroup, error) {
	functionGroup := new(FunctionGroup)
	err := c.db.NewSelect().
		Model(functionGroup).
		Where("id = uuid(?)", groupId).
		Relation("Users").
		Scan(ctx)

	return functionGroup, err
}

func (c *PgConnection) GetEntitledFunctionGroups(userId string, ctx context.Context) ([]string, error) {
	user := new(User)
	err := c.db.NewSelect().
		Model(user).
		Where("id = uuid(?)", userId).
		Relation("FunctionGroups").
		Relation("FunctionGroups.FunctionGroup").
		Scan(ctx)
	functionGroupIds := lo.Map[*FunctionGroupToUserRolePair, string](user.FunctionGroups, func(item *FunctionGroupToUserRolePair, index int) string {
		return item.FunctionGroupId
	})
	return functionGroupIds, err
}

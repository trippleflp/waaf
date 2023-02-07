package postgres

import (
	"context"
	"github.com/uptrace/bun"
)

func (c *pgConnection) IsAdmin(userId, groupId string, ctx context.Context) (bool, error) {
	exists, err := c.db.NewSelect().
		Model((*User)(nil)).
		Where("id = uuid(?)", userId).
		Relation("FunctionGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("function_group_id = uuid(?)", groupId)
		}).
		Exists(ctx)
	return exists, err
}

func (c *pgConnection) FunctionGroupExists(groupId string, ctx context.Context) (bool, error) {
	exists, err := c.db.NewSelect().
		Model((*FunctionGroup)(nil)).
		Where("id = uuid(?)", groupId).
		Exists(ctx)
	return exists, err
}

func (c *pgConnection) GetFunctionGroup(groupId string, ctx context.Context) (*FunctionGroup, error) {
	functionGroup := new(FunctionGroup)
	err := c.db.NewSelect().
		Model(functionGroup).
		Where("id = uuid(?)", groupId).
		Relation("Users").
		Scan(ctx)

	return functionGroup, err
}

package postgres

import (
	"context"
	"github.com/uptrace/bun"
)

func (c *PgConnection) RemoveUsers(userIds []string, functionGroupId string, ctx context.Context) (*FunctionGroup, error) {

	functionGroup := new(FunctionGroup)

	err := c.db.NewSelect().Model(functionGroup).Where("id = uuid(?)", functionGroupId).Relation("Users", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("user_id IN (?)", bun.In(userIds))
	}).Scan(ctx)

	return functionGroup, err
}

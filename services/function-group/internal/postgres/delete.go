package postgres

import (
	"context"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

func (c *PgConnection) RemoveUsers(userIds []string, functionGroupId string, ctx context.Context) (*FunctionGroup, error) {

	functionGroup := new(FunctionGroup)

	err := c.db.NewSelect().Model(functionGroup).Where("id = uuid(?)", functionGroupId).Relation("Users", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("user_id IN (?)", bun.In(userIds))
	}).Relation("Users.User").Scan(ctx)
	if err != nil {
		return nil, err
	}

	pairIds := lo.Map[*FunctionGroupToUserRolePair, int64](functionGroup.Users, func(item *FunctionGroupToUserRolePair, index int) int64 {
		return item.Id
	})

	_, err = c.db.NewDelete().Model((*FunctionGroupToUserRolePair)(nil)).Where(" id IN (?)", bun.In(pairIds)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return c.GetFunctionGroup(functionGroupId, ctx)
}

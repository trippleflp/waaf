package postgres

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
)

func (c *PgConnection) EditUserRole(user *model.UserRolePairInput, functionGroupId string, ctx context.Context) (*FunctionGroup, error) {
	functionGroup := new(FunctionGroup)
	err := c.db.NewSelect().Model(functionGroup).Where("id = uuid(?)", functionGroupId).Relation("Users", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("user_id = uuid(?)", user.UserID)
	}).Relation("Users.User").Scan(ctx)
	if err != nil {
		return nil, err
	}

	_, err = c.db.NewUpdate().Model((*FunctionGroupToUserRolePair)(nil)).Where("id = (?)", functionGroup.Users[0].Id).Set("role = ?", user.Role).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return c.GetFunctionGroup(functionGroupId, ctx)
}

package postgres

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
	"strings"
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

func (c *PgConnection) UpdateFunction(functionTag string, groupId string, ctx context.Context) error {
	functionName := strings.Split(strings.Split(functionTag, "/")[1], ":")[0]
	data := &Function{
		FunctionGroupId: groupId,
		FunctionTag:     functionTag,
		Name:            functionName,
		Id:              fmt.Sprintf("%s/%s", groupId, functionName),
	}

	if _, err := c.db.NewUpdate().
		Model(data).
		Where("id = ?", fmt.Sprintf("%s/%s", groupId, functionName)).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

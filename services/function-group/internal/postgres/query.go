package postgres

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"
	"golang.org/x/crypto/sha3"
)

func (c *PgConnection) IsAdmin(userId, groupId string, ctx context.Context) (bool, error) {
	exists, err := c.db.NewSelect().
		Model((*User)(nil)).
		Where("id = uuid(?)", userId).
		Relation("FunctionGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("role = ?", model.UserRoleAdmin).Where("function_group_id = uuid(?)", groupId)
		}).
		Exists(ctx)
	return exists, err
}

func (c *PgConnection) IsAtLeastReader(userId, groupId string, ctx context.Context) (bool, error) {
	exists, err := c.db.NewSelect().
		Model((*User)(nil)).
		Where("id = uuid(?)", userId).
		Relation("FunctionGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				WhereGroup(" AND  ", func(q *bun.SelectQuery) *bun.SelectQuery {
					return q.Where("role = ?", model.UserRoleAdmin).
						WhereOr("role = ?", model.UserRoleDeveloper).
						WhereOr("role = ?", model.UserRoleReader)
				}).
				Where("function_group_id = uuid(?)", groupId)
		}).
		Exists(ctx)
	return exists, err
}

func (c *PgConnection) IsAtLeastDeveloper(userId, groupId string, ctx context.Context) (bool, error) {
	exists, err := c.db.NewSelect().
		Model((*User)(nil)).
		Where("id = uuid(?)", userId).
		Relation("FunctionGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				WhereGroup(" AND  ", func(q *bun.SelectQuery) *bun.SelectQuery {
					return q.Where("role = ?", model.UserRoleAdmin).
						WhereOr("role = ?", model.UserRoleDeveloper)
				}).
				Where("function_group_id = uuid(?)", groupId)
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
		Relation("Functions").
		//Relation("AllowedFunctionGroups").
		//Relation("AllowedFunctionGroups.ChildFunctionGroup").
		Scan(ctx)

	return functionGroup, err
}

func (c *PgConnection) GetFunctionGroupId(groupName string, ctx context.Context) (string, error) {
	functionGroup := new(FunctionGroup)
	err := c.db.NewSelect().
		Model(functionGroup).
		Where("name = ?", groupName).
		Scan(ctx)

	return functionGroup.Id, err
}

func (c *PgConnection) GetEntitledFunctionGroups(userId string, ctx context.Context) ([]string, error) {
	user := new(User)
	err := c.db.NewSelect().
		Model(user).
		Where("id = uuid(?)", userId).
		Relation("FunctionGroups").
		Relation("FunctionGroups.FunctionGroup").
		Relation("FunctionGroups.FunctionGroup.Functions").
		Scan(ctx)
	functionGroupIds := lo.Map[*FunctionGroupToUserRolePair, string](user.FunctionGroups, func(item *FunctionGroupToUserRolePair, index int) string {
		return item.FunctionGroupId
	})
	return functionGroupIds, err
}

func (c *PgConnection) GetFunctionGroupHash(groupId string, ctx context.Context) (*string, error) {
	functionGroup := new(FunctionGroup)
	err := c.db.NewSelect().
		Model(functionGroup).
		Where("id = uuid(?)", groupId).
		Relation("Users").
		Relation("Functions").
		Scan(ctx)

	bytes, err := json.Marshal(functionGroup)
	if err != nil {
		return nil, err
	}

	h := sha3.New512()
	h.Write(bytes)
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	return &sha1Hash, err
}

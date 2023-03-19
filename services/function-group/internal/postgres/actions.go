package postgres

import (
	"context"
	"fmt"
	"strings"
)

func (c *PgConnection) CreateUserIfNotExist(userId string, ctx context.Context) (bool, error) {
	res, err := c.db.NewInsert().Model(&User{Id: userId}).Ignore().Exec(ctx)
	if err != nil {
		return false, err
	}
	affectedRows, err := res.RowsAffected()
	return affectedRows == 0, err
}

func (c *PgConnection) AddOrUpdateFunction(functionTag string, groupId string, ctx context.Context) error {
	functionName := strings.Split(strings.Split(functionTag, "/")[1], ":")[0]
	exists, err := c.db.NewSelect().
		Model(new(Function)).
		Where("id = ?", fmt.Sprintf("%s/%s", groupId, functionName)).
		Exists(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return c.AddFunction(functionTag, groupId, ctx)
	}
	return c.UpdateFunction(functionTag, groupId, ctx)
}

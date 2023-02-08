package postgres

import (
	"context"
)

func (c *PgConnection) CreateUserIfNotExist(userId string, ctx context.Context) (bool, error) {
	res, err := c.db.NewInsert().Model(&User{Id: userId}).Ignore().Exec(ctx)
	if err != nil {
		return false, err
	}
	affectedRows, err := res.RowsAffected()
	return affectedRows == 0, err
}

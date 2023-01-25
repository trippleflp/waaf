package postgres

import (
	"context"
)

func (c *pgConnection) Exists(userName string, email string, ctx context.Context) (bool, error) {
	exists, err := c.db.NewSelect().Model((*User)(nil)).
		Where("user_name LIKE ?", userName).
		WhereOr("email LIKE ?", email).
		Exists(ctx)

	return exists, err
}

func (c *pgConnection) Authenticate(userName string, email string, passwordHash string, ctx context.Context) (string, error) {
	var user User
	err := c.db.NewSelect().Model((*User)(nil)).
		Where("password = ?", passwordHash).
		Where("email LIKE ?", email).
		WhereOr("user_name LIKE ?", userName).
		Scan(ctx, &user)

	return user.Id, err
}

package postgres

import "context"

func (c *pgConnection) Exists(userName string, email string, ctx context.Context) (bool, error) {
	exists, err := c.db.NewSelect().Model((*User)(nil)).
		Where("user_name LIKE ?", userName).
		WhereOr("email LIKE ?", email).
		Exists(ctx)

	return exists, err
}

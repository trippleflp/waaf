package postgres

import "context"

func (c *pgConnection) CreateUser(user User, ctx context.Context) error {
	_, err := c.db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

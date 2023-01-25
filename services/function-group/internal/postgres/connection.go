package postgres

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"sync"
)

// const connection
type pgConnection struct {
	db *bun.DB
}

var connection = pgConnection{db: nil}
var connectOnce sync.Once

func GetConnection() pgConnection {
	connectOnce.Do(func() {

		sqlDb := sql.OpenDB(pgdriver.NewConnector(
			pgdriver.WithAddr("localhost:5433"),
			pgdriver.WithUser("postgres"),
			pgdriver.WithPassword("postgres"),
			pgdriver.WithInsecure(true)))

		db := bun.NewDB(sqlDb, pgdialect.New())
		connection.db = db
		createSchema(db)
	})
	return connection

}

func createSchema(db *bun.DB) error {
	//res, err := db.NewCreateTable().Model((*User)(nil)).Exec(context.Background())
	err := db.ResetModel(context.Background(), (*User)(nil))
	return err
}

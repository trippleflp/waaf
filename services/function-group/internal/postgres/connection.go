package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
	"sync"
)

// const connection
type PgConnection struct {
	db *bun.DB
}

var connection = PgConnection{db: nil}
var connectOnce sync.Once

func GetConnection() PgConnection {
	connectOnce.Do(func() {
		port := os.Getenv("POSTGRES_PORT")
		if port == "" {
			port = "5432"
		}
		sqlDb := sql.OpenDB(pgdriver.NewConnector(
			pgdriver.WithAddr(fmt.Sprintf("localhost:%s", port)),
			pgdriver.WithUser("postgres"),
			pgdriver.WithPassword("postgres"),
			pgdriver.WithInsecure(true)))

		db := bun.NewDB(sqlDb, pgdialect.New())
		connection.db = db
		err := createSchema(db)
		if err != nil {
			panic(err.Error())
		}
	})
	return connection

}

func createSchema(db *bun.DB) error {
	db.RegisterModel((*FunctionGroupToUserRolePair)(nil))
	models := []interface{}{
		(*Function)(nil),
		(*FunctionGroupToUserRolePair)(nil),
		(*AllowedFunctionGroupPair)(nil),
		(*FunctionGroup)(nil),
		(*User)(nil),
	}

	for _, model := range models {
		if err := db.ResetModel(context.Background(), model); err != nil {
			return err
		}
		//db.RegisterModel(model)
	}
	return nil
}

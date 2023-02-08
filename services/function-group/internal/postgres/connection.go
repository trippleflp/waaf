package postgres

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
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
	//db.RegisterModel((*FunctionGroupToUserRolePair)(nil))
	models := []interface{}{
		(*Function)(nil),
		(*FunctionGroupToUserRolePair)(nil),
		(*FunctionGroup)(nil),
		(*User)(nil),
	}

	for _, model := range models {
		//if err := db.ResetModel(context.Background(), model); err != nil {
		//	return err
		//}
		db.RegisterModel(model)
	}
	//_, err := db.NewCreateTable().Model((*FunctionGroup)(nil)).Exec(context.Background())
	//err := db.ResetModel(context.Background(), (*FunctionGroup)(nil))
	//if err := createSchema(db); err != nil {
	//	return err
	//}

	//values := []interface{}{
	//	&FunctionGroup{Id: "1", Name: "group1"},
	//	&FunctionGroup{Id: "2", Name: "group2"},
	//	&User{Id: "1"},
	//	&FunctionGroupToUserRolePair{UserId: "1", FunctionGroupId: "1", Role: AdminRole},
	//	&FunctionGroupToUserRolePair{UserId: "1", FunctionGroupId: "2", Role: UserRole},
	//}
	//for _, value := range values {
	//	if _, err := db.NewInsert().Model(value).Exec(context.Background()); err != nil {
	//		return err
	//	}
	//}
	//fnGroup := new(FunctionGroup)
	//err := db.NewSelect().
	//	Model(fnGroup).
	//	Relation("Users").
	//	Relation("Users.User").
	//	Limit(1).
	//	Scan(context.Background())
	//if err != nil {
	//	return err
	//}
	//user := new(User)
	//err = db.NewSelect().
	//	Model(user).
	//	Relation("FunctionGroups").
	//	Relation("FunctionGroups.FunctionGroup").
	//	Limit(1).
	//	Scan(context.Background())
	//if err != nil {
	//	return err
	//}
	return nil
}

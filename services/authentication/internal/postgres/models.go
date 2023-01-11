package postgres

type User struct {
	Id       string `bun:",type:uuid"`
	UserName string `bun:",unique"`
	Email    string `bun:",unique"`
	Password string
}

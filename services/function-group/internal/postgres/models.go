package postgres

type Roles string

const (
	AdminRole     Roles = "admin"
	DeveloperRole Roles = "developer"
	ReaderRole    Roles = "reader"
	UserRole      Roles = "user"
)

type RoleUserPair struct {
	role Roles
	user User
}

type User struct {
	Id             string `bun:",type:uuid"`
	FunctionGroups []FunctionGroup
}

type FunctionGroup struct {
	Id    string `bun:",type:uuid"`
	Name  string
	Users []RoleUserPair
	Function
}

type Function struct {
	Id string `bun:",type:uuid"`
}

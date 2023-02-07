package postgres

type Roles string

const (
	AdminRole     Roles = "admin"
	DeveloperRole Roles = "developer"
	ReaderRole    Roles = "reader"
	UserRole      Roles = "user"
)

type FunctionGroupToUserRolePair struct {
	Id              int64 `bun:"id,pk,autoincrement"`
	Role            Roles
	UserId          string
	FunctionGroup   *FunctionGroup `bun:"rel:belongs-to,join:function_group_id=id"`
	User            *User          `bun:"rel:belongs-to,join:user_id=id"`
	FunctionGroupId string
}

type User struct {
	Id             string                         `bun:"type:uuid,pk"`
	FunctionGroups []*FunctionGroupToUserRolePair `bun:"rel:has-many"`
}

type FunctionGroup struct {
	Id          string    `bun:"type:uuid,pk"`
	Name        string    `bun:",unique"`
	functionIds []*string `bun:",array"`
	//Users       []User    `bun:"m2m:function_group_to_user_role_pairs,join:FunctionGroup=User"`
	Users []*FunctionGroupToUserRolePair `bun:"rel:has-many"`
}

type Function struct {
	FunctionId string `bun:",type:uuid,pk"`
}

package postgres

import "gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/model"

type FunctionGroupToUserRolePair struct {
	Id              int64 `bun:"id,pk,autoincrement"`
	Role            model.UserRole
	UserId          string         `bun:"type:uuid"`
	FunctionGroup   *FunctionGroup `bun:"rel:belongs-to,join:function_group_id=id"`
	User            *User          `bun:"rel:belongs-to,join:user_id=id"`
	FunctionGroupId string         `bun:"type:uuid"`
}

type AllowedFunctionGroupPair struct {
	Id                    int64          `bun:"id,pk,autoincrement"`
	ParentFunctionGroup   *FunctionGroup `bun:"rel:belongs-to,join:parent_function_group_id=id"`
	ParentFunctionGroupId string         `bun:"type:uuid"`
	ChildFunctionGroup    *FunctionGroup `bun:"rel:belongs-to,join:child_function_group_id=id"`
	ChildFunctionGroupId  string         `bun:"type:uuid"`
}

type User struct {
	Id             string                         `bun:"type:uuid,pk"`
	FunctionGroups []*FunctionGroupToUserRolePair `bun:"rel:has-many"`
}

type FunctionGroup struct {
	Id                    string                         `bun:"type:uuid,pk"`
	Name                  string                         `bun:",unique"`
	Functions             []*Function                    `bun:"rel:has-many"`
	Users                 []*FunctionGroupToUserRolePair `bun:"rel:has-many"`
	AllowedFunctionGroups []*AllowedFunctionGroupPair    `bun:"rel:has-many"`
}

type Function struct {
	FunctionId      string `bun:",type:uuid,pk"`
	Name            string `bun:",unique"`
	FunctionGroupId string `bun:"type:uuid"`
}

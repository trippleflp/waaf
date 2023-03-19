package postgres

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

var _ = []bun.BeforeAppendModelHook{
	(*Function)(nil),
	(*FunctionGroupToUserRolePair)(nil),
	(*AllowedFunctionGroupPair)(nil),
	(*FunctionGroup)(nil),
}

func (u *Function) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.LastChanged = time.Now()
	case *bun.UpdateQuery:
		u.LastChanged = time.Now()
	}
	return nil
}
func (u *FunctionGroupToUserRolePair) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.LastChanged = time.Now()
	case *bun.UpdateQuery:
		u.LastChanged = time.Now()
	}
	return nil
}
func (u *AllowedFunctionGroupPair) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.LastChanged = time.Now()
	case *bun.UpdateQuery:
		u.LastChanged = time.Now()
	}
	return nil
}
func (u *FunctionGroup) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.LastChanged = time.Now()
	case *bun.UpdateQuery:
		u.LastChanged = time.Now()
	}
	return nil
}

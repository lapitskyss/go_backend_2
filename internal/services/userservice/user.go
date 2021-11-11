package userservice

import (
	"context"
)

type User struct {
	ID   uint64
	Name string
}

type NewUser struct {
	Name string
}

type UserWithScopes struct {
	User
	Scopes []string
}

type UsersStore interface {
	Create(ctx context.Context, u NewUser) (*User, error)
	AddUserToScope(ctx context.Context, userId uint64, scopeId uint64) (*UserWithScopes, error)
	RemoveUserFromScope(ctx context.Context, userId uint64, scopeId uint64) (*UserWithScopes, error)
	SearchUser(ctx context.Context, query string) ([]User, error)
}

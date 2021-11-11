package userstore

import (
	"context"

	"github.com/lapitskyss/go_backend_2/internal/services/userservice"
)

type UserStore struct {
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (s *UserStore) Create(ctx context.Context, u userservice.NewUser) (*userservice.User, error) {
	panic("implement me")
}

func (s *UserStore) AddUserToScope(ctx context.Context, userId uint64, scopeId uint64) (*userservice.UserWithScopes, error) {
	panic("implement me")
}

func (s *UserStore) RemoveUserFromScope(ctx context.Context, userId uint64, scopeId uint64) (*userservice.UserWithScopes, error) {
	panic("implement me")
}

func (s *UserStore) SearchUser(ctx context.Context, query string) ([]userservice.User, error) {
	panic("implement me")
}

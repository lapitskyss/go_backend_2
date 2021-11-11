package userservice

import "context"

type CreateUserRequest struct {
	Name string
}

func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	return s.UserStore.Create(ctx, NewUser{
		Name: req.Name,
	})
}

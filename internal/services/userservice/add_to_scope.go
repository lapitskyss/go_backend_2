package userservice

import "context"

type AddUserToScopeRequest struct {
	UserID  uint64
	ScopeID uint64
}

type AddUserToScopeResponse struct {
	ID     uint64
	Name   string
	Scopes []string
}

func (s *Service) AddUserToScope(ctx context.Context, req *AddUserToScopeRequest) (*AddUserToScopeResponse, error) {
	_, _ = s.UserStore.AddUserToScope(ctx, req.UserID, req.ScopeID)

	return nil, nil
}

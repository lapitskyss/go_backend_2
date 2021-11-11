package userservice

import "context"

type RemoveUserFromScopeRequest struct {
	UserID  uint64
	ScopeID uint64
}

type RemoveUserFromScopeResponse struct {
	ID     uint64
	Name   string
	Scopes []string
}

func (s *Service) RemoveUserFromScope(ctx context.Context, req *RemoveUserFromScopeRequest) (*RemoveUserFromScopeResponse, error) {
	_, _ = s.UserStore.RemoveUserFromScope(ctx, req.UserID, req.ScopeID)

	return nil, nil
}

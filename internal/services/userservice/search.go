package userservice

import "context"

type SearchUsersRequest struct {
	query string
}

func (s *Service) SearchUsers(ctx context.Context, req *SearchUsersRequest) ([]User, error) {
	return s.UserStore.SearchUser(ctx, req.query)
}

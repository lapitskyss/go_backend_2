package scopeservice

import "context"

type SearchScopesRequest struct {
	query string
}

func (s *Service) SearchScopes(ctx context.Context, req *SearchScopesRequest) ([]Scope, error) {
	return s.ScopeStore.SearchScope(ctx, req.query)
}

package scopeservice

import "context"

type CreateScopeRequest struct {
	Name string
	Type Type
}

func (s *Service) CreateScope(ctx context.Context, req *CreateScopeRequest) (*Scope, error) {
	return s.ScopeStore.Create(ctx, NewScope{
		Name: req.Name,
		Type: req.Type,
	})
}

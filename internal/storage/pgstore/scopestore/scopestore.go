package scopestore

import (
	"context"

	"github.com/lapitskyss/go_backend_2/internal/services/scopeservice"
)

type ScopeStore struct {
}

func NewScopeStore() *ScopeStore {
	return &ScopeStore{}
}

func (s *ScopeStore) Create(ctx context.Context, scope scopeservice.NewScope) (*scopeservice.Scope, error) {
	panic("implement me")
}

func (s *ScopeStore) SearchScope(ctx context.Context, query string) ([]scopeservice.Scope, error) {
	panic("implement me")
}

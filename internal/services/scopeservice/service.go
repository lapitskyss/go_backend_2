package scopeservice

type Service struct {
	ScopeStore ScopesStore
}

func NewScopeService(scopeStore ScopesStore) *Service {
	return &Service{
		ScopeStore: scopeStore,
	}
}

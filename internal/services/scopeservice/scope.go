package scopeservice

import (
	"context"
)

type Scope struct {
	ID   uint64
	Name string
	Type Type
}

type Type uint64

const (
	ProjectType Type = iota
	OrganizationType
	CorporateGroupType
	CommunityType
)

type NewScope struct {
	Name string
	Type Type
}

type ScopesStore interface {
	Create(ctx context.Context, scope NewScope) (*Scope, error)
	SearchScope(ctx context.Context, query string) ([]Scope, error)
}

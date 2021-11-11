package pgstore

import (
	"context"
)

type Store struct {
}

func Connect(ctx context.Context, url string) (*Store, error) {
	return nil, nil
}

func (s *Store) Close() {
	return
}

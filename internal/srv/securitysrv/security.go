package securitysrv

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"go.uber.org/zap"
)

type SecuritySrv struct {
	store  SecurityStore
	random *rand.Rand
	log    *zap.Logger
}

func InitSecuritySrv(store SecurityStore, log *zap.Logger) *SecuritySrv {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return &SecuritySrv{
		store:  store,
		random: r,
		log:    log,
	}
}

var (
	ErrCodeNotFound = errors.New("code not found")
)

type SecurityStore interface {
	SaveCode(ctx context.Context, key string, value string) error
	GetCode(ctx context.Context, key string) (string, error)
}

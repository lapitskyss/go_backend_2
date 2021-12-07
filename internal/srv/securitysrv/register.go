package securitysrv

import (
	"context"
	"errors"
	"math/rand"
	"strconv"

	"go.uber.org/zap"

	"github.com/lapitskyss/go_backend_2/internal/pkg/response"
)

const (
	stepCode     = "code"
	stepRegister = "register"
)

var (
	ErrIncorrectConfirmCode = errors.New("incorrect confirm code")
	ErrExpiredConfirmCode   = errors.New("confirm code expired")
)

type RegisterRequest struct {
	Step     string // code/register
	Email    string
	Password string
	Code     string
}

type RegisterResponse struct {
	Message string
}

func (s *SecuritySrv) RegisterEmail(ctx context.Context, r RegisterRequest) (*RegisterResponse, error) {
	// TODO: validate request parameters

	if r.Step == stepCode {
		err := s.sendEmailCode(ctx, r)
		if err != nil {
			return nil, err
		}

		return &RegisterResponse{
			Message: "Code send to email address",
		}, nil
	}

	err := s.registerUser(ctx, r)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{
		Message: "Success register",
	}, nil
}

func (s *SecuritySrv) sendEmailCode(ctx context.Context, r RegisterRequest) error {
	// TODO: add spam detection
	key := "email:" + r.Email
	code := s.getRandomCode()

	// ... sending code to user
	s.log.Info("confirm code: " + code)

	return s.store.SaveCode(ctx, key, code)
}

func (s *SecuritySrv) registerUser(ctx context.Context, r RegisterRequest) error {
	key := "email:" + r.Email
	code, err := s.store.GetCode(ctx, key)
	if err != nil {
		if errors.Is(err, ErrCodeNotFound) {
			return response.ErrBadRequest(ErrExpiredConfirmCode)
		}

		s.log.Error(" Err registerUser GetCode", zap.Error(err))

		return response.ErrInternal()
	}

	s.log.Info("entered code: " + r.Code)
	s.log.Info("required code: " + code)

	if code != r.Code {
		// TODO: add limits for check code

		return response.ErrBadRequest(ErrIncorrectConfirmCode)
	}
	// ... register user

	return nil
}

func (s *SecuritySrv) getRandomCode() string {
	code := 1000 + rand.Intn(9999-1000)

	return strconv.Itoa(code)
}

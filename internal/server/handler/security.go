package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/lapitskyss/go_backend_2/internal/pkg/render"
	"github.com/lapitskyss/go_backend_2/internal/pkg/response"
	"github.com/lapitskyss/go_backend_2/internal/srv/securitysrv"
)

type Handler struct {
	security *securitysrv.SecuritySrv
	log      *zap.Logger
}

func InitHandler(security *securitysrv.SecuritySrv, log *zap.Logger) *Handler {
	return &Handler{
		security: security,
		log:      log,
	}
}

type registerUserRequest struct {
	Step     string `json:"step"`
	Email    string `email:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type registerUserResponse struct {
	Message string `json:"message"`
}

func (h *Handler) RegisterEmail(w http.ResponseWriter, r *http.Request) {
	var req = &registerUserRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		render.BadRequestError(w, errors.New("incorrect request params"))
		return
	}

	result, err := h.security.RegisterEmail(r.Context(), securitysrv.RegisterRequest{
		Step:     req.Step,
		Email:    req.Email,
		Password: req.Password,
		Code:     req.Code,
	})
	if err != nil {
		response.SendErrorResponse(w, err)
		return
	}

	render.Success(w, registerUserResponse{
		Message: result.Message,
	})
}

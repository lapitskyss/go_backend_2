package handler

import (
	"net/http"

	"github.com/lapitskyss/go_backend_2/internal/services/userservice"
)

type UserHandler struct {
	UserService *userservice.Service
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	//....
	_, _ = h.UserService.CreateUser(r.Context(), &userservice.CreateUserRequest{})
	//....
}

func (h *UserHandler) AddUserToScope(w http.ResponseWriter, r *http.Request) {
	//....
	_, _ = h.UserService.AddUserToScope(r.Context(), &userservice.AddUserToScopeRequest{})
	//....
}

func (h *UserHandler) RemoveUserFromScope(w http.ResponseWriter, r *http.Request) {
	//....
	_, _ = h.UserService.RemoveUserFromScope(r.Context(), &userservice.RemoveUserFromScopeRequest{})
	//....
}

func (h *UserHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	//....
	_, _ = h.UserService.SearchUsers(r.Context(), &userservice.SearchUsersRequest{})
	//....
}

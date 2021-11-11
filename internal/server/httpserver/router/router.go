package router

import (
	"net/http"

	"github.com/lapitskyss/go_backend_2/internal/server/httpserver/handler"
	"github.com/lapitskyss/go_backend_2/internal/services/scopeservice"
	"github.com/lapitskyss/go_backend_2/internal/services/userservice"
)

func NewRouter(userService *userservice.Service, scopeService *scopeservice.Service) http.Handler {
	uh := handler.UserHandler{UserService: userService}
	sh := handler.ScopeHandler{ScopeService: scopeService}

	r := http.NewServeMux()

	r.Handle("/user/create", http.HandlerFunc(uh.CreateUser))
	r.Handle("/user/scope/add", http.HandlerFunc(uh.AddUserToScope))
	r.Handle("/user/scope/remove", http.HandlerFunc(uh.RemoveUserFromScope))
	r.Handle("/user/search", http.HandlerFunc(uh.SearchUsers))

	r.Handle("/scope/create", http.HandlerFunc(sh.CreateScope))
	r.Handle("/scope/search", http.HandlerFunc(sh.SearchScopes))

	return r
}

package handler

import (
	"net/http"

	"github.com/lapitskyss/go_backend_2/internal/services/scopeservice"
)

type ScopeHandler struct {
	ScopeService *scopeservice.Service
}

func (h *ScopeHandler) CreateScope(w http.ResponseWriter, r *http.Request) {
	//....
	_, _ = h.ScopeService.CreateScope(r.Context(), &scopeservice.CreateScopeRequest{})
	//....
}

func (h *ScopeHandler) SearchScopes(w http.ResponseWriter, r *http.Request) {
	//....
	_, _ = h.ScopeService.SearchScopes(r.Context(), &scopeservice.SearchScopesRequest{})
	//....
}

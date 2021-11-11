package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/lapitskyss/go_backend_2/internal/server/httpserver/router"
	"github.com/lapitskyss/go_backend_2/internal/services/scopeservice"
	"github.com/lapitskyss/go_backend_2/internal/services/userservice"
)

type HTTPServer struct {
	server http.Server
	errors chan error
}

func NewHTTPServer(userService *userservice.Service, scopeService *scopeservice.Service) *HTTPServer {
	r := router.NewRouter(userService, scopeService)

	return &HTTPServer{
		server: http.Server{
			Addr:    ":3000",
			Handler: r,

			ReadTimeout:       1 * time.Second,
			WriteTimeout:      90 * time.Second,
			IdleTimeout:       30 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
		},
	}
}

func (srv *HTTPServer) Start() {
	go func() {
		err := srv.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			srv.errors <- err
			close(srv.errors)
		}
	}()
}

func (srv *HTTPServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.server.Shutdown(ctx)
}

func (srv *HTTPServer) Notify() <-chan error {
	return srv.errors
}

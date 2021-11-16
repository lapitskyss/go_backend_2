package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type RESTServer struct {
	server http.Server
	errors chan error
}

func NewRESTServer() *RESTServer {
	r := chi.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsHandler.Handler)

	r.Post("/v1/hello", helloHandler)
	r.Get("/__heartbeat__", heartbeatHandler)

	return &RESTServer{
		server: http.Server{
			Addr:    ":3000",
			Handler: r,
		},
	}
}

func (s *RESTServer) Start() {
	go func() {
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.errors <- err
			close(s.errors)
		}
	}()
}

func (s *RESTServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func (s *RESTServer) Notify() <-chan error {
	return s.errors
}

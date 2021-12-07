package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"

	"github.com/lapitskyss/go_backend_2/internal/server/handler"
)

type Server struct {
	server http.Server
	log    *zap.Logger
	errors chan error
}

func InitServer(port string, handler *handler.Handler, log *zap.Logger) *Server {
	r := chi.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(middleware.Recoverer)
	r.Use(corsHandler.Handler)
	r.Use(middleware.AllowContentType("application/json"))

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/register/email", handler.RegisterEmail)
	})

	return &Server{
		server: http.Server{
			Addr:    ":" + port,
			Handler: r,

			ReadTimeout:       1 * time.Second,
			WriteTimeout:      90 * time.Second,
			IdleTimeout:       30 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
		},
		log: log,
	}
}

func (s *Server) Start() {
	s.log.Info("Server started on port " + s.server.Addr + ".")
	go func() {
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.log.Error("Server return error", zap.Error(err))
			s.errors <- err
		}
	}()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.errors
}

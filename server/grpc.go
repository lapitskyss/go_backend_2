package server

import (
	"net"

	"google.golang.org/grpc"

	pb "github.com/lapitskyss/go_backend_2/proto"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type GRPCServer struct {
	server *grpc.Server
	errors chan error
}

func NewGRPCServer() *GRPCServer {
	srv := grpc.NewServer()
	svc := &server{}

	pb.RegisterGreeterServer(srv, svc)
	healthpb.RegisterHealthServer(srv, svc)

	return &GRPCServer{
		server: srv,
	}
}

func (s *GRPCServer) Start() {
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			s.errors <- err
			close(s.errors)
			return
		}
		err = s.server.Serve(lis)
		if err != nil {
			s.errors <- err
			close(s.errors)
		}
	}()
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}

func (s *GRPCServer) Notify() <-chan error {
	return s.errors
}

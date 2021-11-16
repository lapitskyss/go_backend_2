package server

import (
	"context"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/lapitskyss/go_backend_2/proto"
	"github.com/lapitskyss/go_backend_2/server/greeter"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Check(ctx context.Context, request *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	panic("implement me")
}

func (s *server) Watch(request *healthpb.HealthCheckRequest, watchServer healthpb.Health_WatchServer) error {
	panic("implement me")
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	message := greeter.SayHello(ctx, in.GetName())

	return &pb.HelloReply{Message: message}, nil
}

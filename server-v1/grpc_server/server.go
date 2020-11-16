package grpc_server

import (
	"context"

	"github.com/bygui86/go-grpc-compatibility/server-v1/domain"
	"github.com/bygui86/go-grpc-compatibility/server-v1/logger"
)

// Server implements helloworld.GreeterServer
type Server struct{}

// SayHello implements service HelloService helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *domain.HelloRequest) (*domain.HelloResponse, error) {
	logger.SugaredLogger.Info("Say hello to world")
	return &domain.HelloResponse{
		Greeting: "Hello, world!",
	}, nil
}

package grpc_server

import (
	"context"
	"fmt"

	"github.com/bygui86/go-grpc-compatibility/server-v2/domain"
	"github.com/bygui86/go-grpc-compatibility/server-v2/logger"
)

const (
	emptyName = "<EMPTY>"
)

// Server implements helloworld.GreeterServer
type Server struct{}

// SayHello implements service HelloService helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *domain.HelloRequest) (*domain.HelloResponse, error) {
	var name string
	if in.Name != "" {
		name = in.Name
	} else {
		name = emptyName
	}
	logger.SugaredLogger.Infof("Say hello to %s", name)

	return &domain.HelloResponse{
		Greeting:      "Hello!",
		HiddenMessage: buildHiddenMsg(name),
	}, nil
}

func buildHiddenMsg(name string) string {
	return fmt.Sprintf("This is a message reserved to %s", name)
}

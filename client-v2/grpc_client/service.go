package grpc_client

import (
	"google.golang.org/grpc"
)

type GrpcGreetingService struct {
	GrpcConnection *grpc.ClientConn
}

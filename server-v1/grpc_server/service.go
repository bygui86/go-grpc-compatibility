package grpc_server

import "net"

type GrpcHelloService struct {
	Listener net.Listener
	Network  string
	Address  string
}

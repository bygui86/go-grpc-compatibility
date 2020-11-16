package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bygui86/go-grpc-compatibility/server-v1/domain"
	"github.com/bygui86/go-grpc-compatibility/server-v1/grpc_server"
	"github.com/bygui86/go-grpc-compatibility/server-v1/logger"
	"github.com/bygui86/go-grpc-compatibility/server-v1/utils"

	"google.golang.org/grpc"
)

const (
	serverAddressEnvVar = "GRPC_SERVER_ADDRESS"

	serverAddressEnvVarDefault = "0.0.0.0:50051"

	grpcListenerNetwork = "tcp"
)

func main() {
	address := utils.GetString(serverAddressEnvVar, serverAddressEnvVarDefault)

	listener := createListener(grpcListenerNetwork, address)
	defer listener.Close()
	logger.SugaredLogger.Infof("TCP listener ready on %s", address)

	go startGrpcServer(listener)
	logger.SugaredLogger.Infof("gRPC server ready")

	logger.SugaredLogger.Info("Hello service started!")
	startSysCallChannel()
}

func createListener(network, address string) net.Listener {
	listener, err := net.Listen(network, address)
	if err != nil {
		logger.SugaredLogger.Errorf("Failed to listen: %v", err.Error())
		os.Exit(3)
	}

	return listener
}

func startGrpcServer(listener net.Listener) {
	grpcServer := grpc.NewServer()
	helloSvcServer := grpc_server.Server{}
	domain.RegisterHelloServiceServer(grpcServer, &helloSvcServer)
	err := grpcServer.Serve(listener)
	if err != nil {
		logger.SugaredLogger.Errorf("Failed to serve: %v", err)
		os.Exit(4)
	}
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	logger.SugaredLogger.Info("Termination signal received!")
}

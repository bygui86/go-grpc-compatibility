package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-grpc-compatibility/client-v1/domain"
	"github.com/bygui86/go-grpc-compatibility/client-v1/logger"
	"github.com/bygui86/go-grpc-compatibility/client-v1/utils"

	"google.golang.org/grpc"
)

const (
	serverAddressEnvVar = "SERVER_ADDRESS"

	serverAddressEnvVarDefault = "0.0.0.0:50051"
)

func main() {
	serverAddress := utils.GetString(serverAddressEnvVar, serverAddressEnvVarDefault)

	grpcConn := createGrpcConnection(serverAddress)
	defer grpcConn.Close()
	logger.SugaredLogger.Infof("gRPC Connection ready to %s", serverAddress)

	go startGreetings(grpcConn)

	logger.SugaredLogger.Info("Greeting service started!")
	startSysCallChannel()
}

func createGrpcConnection(host string) *grpc.ClientConn {
	connection, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		logger.SugaredLogger.Errorf("Connection to gRPC server failed: %v", err.Error())
		os.Exit(3)
	}

	logger.SugaredLogger.Info("State: ", connection.GetState())
	logger.SugaredLogger.Info("Target: ", connection.Target())

	return connection
}

func startGreetings(connection *grpc.ClientConn) {
	timeout := 2 * time.Second
	client := domain.NewHelloServiceClient(connection)
	logger.SugaredLogger.Info("Starting greeting the world...")
	for {
		go greet(client, timeout)
		time.Sleep(3 * time.Second)
	}
}

func greet(client domain.HelloServiceClient, timeout time.Duration) {
	// WARNING: the connection context is one-shot, it must be refreshed before every request
	connectionCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	response, err := client.SayHello(connectionCtx, &domain.HelloRequest{})
	if err != nil {
		logger.SugaredLogger.Errorf("Could not greet world: %v", err.Error())
		return
	}
	logger.SugaredLogger.Infof("Greeting: %s", response.Greeting)
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	logger.SugaredLogger.Info("Termination signal received!")
}

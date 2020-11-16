package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-grpc-compatibility/client-v2/domain"
	"github.com/bygui86/go-grpc-compatibility/client-v2/logger"
	"github.com/bygui86/go-grpc-compatibility/client-v2/utils"

	"google.golang.org/grpc"
)

const (
	serverAddressEnvVar = "SERVER_ADDRESS"
	greeetingNameEnvVar = "GREETING_NAME"

	serverAddressEnvVarDefault = "0.0.0.0:50051"
	greeetingNameEnvVarDefault = "ANONYMOUS"
)

func main() {
	serverAddress := utils.GetString(serverAddressEnvVar, serverAddressEnvVarDefault)
	name := utils.GetString(greeetingNameEnvVar, greeetingNameEnvVarDefault)

	grpcConn := createGrpcConnection(serverAddress)
	defer grpcConn.Close()
	logger.SugaredLogger.Infof("gRPC Connection ready to %s", serverAddress)

	go startGreetings(grpcConn, name)

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

func startGreetings(connection *grpc.ClientConn, name string) {
	timeout := 2 * time.Second
	client := domain.NewHelloServiceClient(connection)
	logger.SugaredLogger.Infof("Starting greeting %s...", name)
	for {
		go greet(client, timeout, name)
		time.Sleep(3 * time.Second)
	}
}

func greet(client domain.HelloServiceClient, timeout time.Duration, name string) {
	// WARNING: the connection context is one-shot, it must be refreshed before every request
	connectionCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	response, err := client.SayHello(connectionCtx, &domain.HelloRequest{Name: name})
	if err != nil {
		logger.SugaredLogger.Errorf("Could not greet %s: %v", name, err.Error())
		return
	}
	logger.SugaredLogger.Infof("Greeting: %s - HiddenMessage: %s", response.Greeting, response.HiddenMessage)
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	logger.SugaredLogger.Info("Termination signal received!")
}

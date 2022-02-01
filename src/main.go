package main

import (
	"fmt"
	userpb "github.com/asazorojas/go-grpc-users-api/common/protos/v1/user"
	"github.com/asazorojas/go-grpc-users-api/src/server"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

const portNumber = "9000"
const network = "tcp"

func main() {
	listener, err := net.Listen(network, fmt.Sprintf(":%s", portNumber))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		)),
	)
	userpb.RegisterUserServer(grpcServer, &server.UserServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

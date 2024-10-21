package main

import (
	"log"
	"net"

	pb "greeter-grpc/api/proto/v1"

	"google.golang.org/grpc"

	"greeter-grpc/internal/services"
)

func main() {
	listener, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	service := &services.GreeterServiceServer{}

	pb.RegisterGreeterServiceServer(grpcServer, service)
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("error starting server %s", err)
	}
}
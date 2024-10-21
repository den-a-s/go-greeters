package services

import (
	"context"
	"fmt"
	pb "greeter-grpc/api/proto/v1"
)

type GreeterServiceServer struct {
	pb.UnimplementedGreeterServiceServer
}

func (service *GreeterServiceServer) Greet(_ context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Got name: %s\n", req.GetName())
	fmt.Printf("Hello, %s\n", req.GetName())
	return &pb.GreetResponse{}, nil
}

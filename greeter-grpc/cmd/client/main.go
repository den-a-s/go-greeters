package main

import (
	"context"
	"log"

	pb "greeter-grpc/api/proto/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient(":8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("init conn: %s", err)
	}

	defer conn.Close()

	client := pb.NewGreeterServiceClient(conn)

	resp, err := client.Greet(context.Background(), &pb.GreetRequest{Name: "Dennis"})
	if err != nil {
		log.Fatalf("error on query: %s", err)
	}

	log.Println(resp)
}
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/PJ9172/Go-Lang-Programming/gRPC/hello_gRPC/proto"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewGreeterClient(conn)

	// Prepare the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the RPC call
	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Go Developer"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	// Print the response
	log.Printf("Server says: %s", response.Message)
}

// server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/PJ9172/Go-Lang-Programming/gRPC/hello_gRPC/proto"

	"google.golang.org/grpc"
)

// Server struct implements the GreeterServer interface
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements the SayHello RPC
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received request from: %s", req.Name)
	return &pb.HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	fmt.Println("gRPC server listening on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

package main

import (
	"Unary_RPC/greetpb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func main() {
	fmt.Println("Hello I'm Server!!!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to Listen : %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	log.Println("Server is running on port : 50051")
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}

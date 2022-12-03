package main

import (
	"fmt"
	"log"
	"net"

	"grpc_example/hello"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("grpc")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %+v", err)
	}
	grpcServer := grpc.NewServer()
	s := hello.Server{}

	hello.RegisterHelloServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %+v", err)
	}

	// Go func ile ayrı bir portta başka bir grpc server aç
}

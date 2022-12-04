package main

import (
	"context"
	"log"

	"grpc_example/hello"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %+v", err)
	}
	defer conn.Close()

	c := hello.NewHelloServiceClient(conn)
	m, err := c.Talkie(context.Background(), &hello.Message{Content: "from client"})
	if err != nil {
		log.Fatalf("Failed: %+v", err)
	}
	log.Printf(m.Content)
}

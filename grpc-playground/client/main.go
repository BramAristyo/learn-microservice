package main

import (
	"context"
	"log"

	pb "github.com/bramAristyo/learn-microservice/grpc-playground"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	cc, err := grpc.NewClient(":50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	c := pb.NewGreeterServiceClient(cc)

	var callOpts []grpc.CallOption
	log.Printf("Request to GreeterService SayHello('John Doe')\n")

	ctx := metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs(
			"authorization", "Bearer EXAMPLE",
			"x-request-id", "request-key-001",
		),
	)

	responseHello, err := c.SayHello(ctx, &pb.SayHelloRequest{Name: "John Doe"}, callOpts...)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", responseHello.Message)
}

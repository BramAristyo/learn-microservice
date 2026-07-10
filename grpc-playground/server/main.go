package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/bramAristyo/learn-microservice/grpc-playground"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	time.Sleep(3 * time.Second)
	return &pb.SayHelloResponse{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50050))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, &Server{})

	go func() {
		fmt.Printf("server run on :%d\n", 50050)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.GracefulStop()
	log.Println("bye..")
}

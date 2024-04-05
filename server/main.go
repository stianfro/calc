package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/stianfro/calc/gen/go/calculator/v1"
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Print("received add request")
	return &pb.AddResponse{
		Result: in.A + in.B,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to serve:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve:", err)
	}
}

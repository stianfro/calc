package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

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

func (s *server) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideResponse, error) {
	log.Print("received divide request")

	if in.B == 0 {
		return nil, status.Error(
			codes.InvalidArgument, "Cannot device by zero",
		)
	}

	return &pb.DivideResponse{
		Result: in.A / in.B,
	}, nil
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Print("received sum request")

	var sum int64

	for _, number := range in.Numbers {
		sum += number
	}

	return &pb.SumResponse{
		Result: sum,
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

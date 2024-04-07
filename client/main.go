package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/stianfro/calc/gen/go/calculator/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	serverAddress := flag.String(
		"server", "localhost:8080",
		"The server address in the format host:port",
	)
	flag.Parse()

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(*serverAddress, opts...)
	if err != nil {
		log.Fatalln("fail to dial:", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	res, err := client.Sum(ctx, &pb.SumRequest{
		Numbers: []int64{10, 10, 10, 10, 10},
	})
	if err != nil {
		log.Fatalln("error sending request:", err)
	}

	fmt.Println("result:", res.Result)
}

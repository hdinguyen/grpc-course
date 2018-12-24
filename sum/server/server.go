package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hdinguyen/grpc-course/sum/sumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Sum function invoke: %v", req)
	result := req.GetNumber().GetNum1() + req.GetNumber().GetNum2()
	res := &sumpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Sum server start")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

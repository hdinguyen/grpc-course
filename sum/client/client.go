package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hdinguyen/grpc-course/sum/sumpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Sum client start")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(" Client dial err: %v", err)
	}

	defer cc.Close()
	c := sumpb.NewSumServiceClient(cc)

	req := &sumpb.SumRequest{
		Number: &sumpb.Sum{
			Num1: 3,
			Num2: 4,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling: %v", err)
	}
	log.Printf(" Sum service response: %v", res)
}

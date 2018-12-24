package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hdinguyen/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "ABC",
			LastName:  "Maarek",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling: %v", err)
	}

	log.Printf(" Response from Greet: %v", res)
}

package main

import (
	"context"
	"fmt"
	"grpc-calculator/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing to the server %v\n", err)
	}
	client := proto.NewCalculatorServiceClient(conn)
	addRequest := &proto.CalculatorRequest{X: 100, Y: 200}
	addResult, err := client.Add(context.Background(), addRequest)
	if err != nil {
		log.Fatalf("Error Adding numbers %v\n", err)
	}
	fmt.Printf("Adding %d and %d results in %d \n", addRequest.X, addRequest.Y, addResult.GetResult())
}

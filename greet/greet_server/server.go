package main

//My name is mihirkelkar.

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mihirkelkar/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

//A new interface got defined based on our .proto file's functions in the
//.pb.go file. We need to implement the interface here.

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	greeting := req.GetGreeting()
	firstName := greeting.GetFirstName()
	lastName := greeting.GetLastName()

	res := "Hello " + firstName + lastName
	return &greetpb.GreetResponse{
		Result: res,
	}, nil
}

func main() {
	fmt.Println("Hello World from Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listed %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

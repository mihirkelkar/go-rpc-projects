package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mihirkelkar/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Add(ctx context.Context, sp *greetpb.SumRequest) (*greetpb.SumResponse, error) {
	ans := &greetpb.SumResponse{}
	ans.Answer = 0
	add := sp.GetAddition()
	ans.Answer = add.GetFirstparam() + add.GetSecondparam()
	return ans, nil
}

func main() {
	fmt.Println("Hello from the Addition Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listed %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterAdditionServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

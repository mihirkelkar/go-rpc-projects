package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mihirkelkar/grpc/primenum/primepb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Factorize(req *primepb.PrimeRequest, stream primepb.PrimeService_FactorizeServer) error {
	input := int(req.Message.Number)
	k := 2
	for input > 1 {
		if input%k == 0 {
			resp := &primepb.PrimeResponse{Number: int32(k)}
			stream.Send(resp)
			time.Sleep(1000 * time.Millisecond)
			input = input / k
		} else {
			k = k + 1
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello World from Streaming GRPC Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listed %v", err)
	}

	s := grpc.NewServer()
	primepb.RegisterPrimeServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

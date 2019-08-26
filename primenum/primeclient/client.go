package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/mihirkelkar/grpc/primenum/primepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from Client !")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()

	client := primepb.NewPrimeServiceClient(conn)

	req := &primepb.PrimeRequest{Message: &primepb.PrimeMessage{Number: 10}}

	rescStream, err := client.Factorize(context.Background(), req)
	if err != nil {
		log.Fatalf("Reponse from server returned error, %v", err)
	}

	for {
		msg, err := rescStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("err while parsing response, %v", err)
		}

		log.Printf("Factor from Stream %v", msg.GetNumber())
	}

}

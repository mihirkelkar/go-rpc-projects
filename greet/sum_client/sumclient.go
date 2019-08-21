package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mihirkelkar/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from Client !")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()

	client := greetpb.NewAdditionServiceClient(conn)

	req := &greetpb.SumRequest{
		Addition: &greetpb.Addition{
			Firstparam:  11,
			Secondparam: 13,
		},
	}

	resp, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Reponse from server returned error, %v", err)
	}

	fmt.Println(resp)
	return

}

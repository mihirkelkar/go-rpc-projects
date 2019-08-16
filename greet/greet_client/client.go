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

	client := greetpb.NewGreetServiceClient(conn)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mihir",
			LastName:  "Kelkar",
		},
	}

	resp, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Reponse from server returned error, %v", err)
	}

	fmt.Println(resp)
	return

}

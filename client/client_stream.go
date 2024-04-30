package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SaiSawant1/go-grpc/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("Client Stream started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send name %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while Sending %v", err)
		}
		log.Printf("Sent the request with name %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Stream finished")
	if err != nil {
		log.Fatalf("Errow while Receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}

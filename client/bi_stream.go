package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/SaiSawant1/go-grpc/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client Side bi-directional streaming started")
	stream, err := client.SayHelloBidirectionsStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names")
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}

			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending names: %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}

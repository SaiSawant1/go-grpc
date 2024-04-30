package main

import (
	"io"
	"log"

	pb "github.com/SaiSawant1/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionsStreaming(stream pb.GreetService_SayHelloBidirectionsStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got Request with the name %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

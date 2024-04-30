package main

import (
	"context"

	pb "github.com/SaiSawant1/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, request *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

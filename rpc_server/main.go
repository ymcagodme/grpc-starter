package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ymcagodme/shortn/proto"
	"google.golang.org/grpc"
)

const (
	port = ":52002"
)

type server struct {
	pb.UnimplementedShortnServer
}

func (s *server) AddPageRpc(ctx context.Context, in *pb.AddPageRequest) (*pb.AddPageResponse, error) {
	log.Printf("AddPageRequest: incoming raw_url = %s", in.GetRawUrl())
	return &pb.AddPageResponse{ShortUrl: "this is short url"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen to %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterShortnServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

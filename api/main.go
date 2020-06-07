package main

import (
	"context"
	"log"
	"net"

	"github.com/ymcagodme/shortn/api/core"
	pb "github.com/ymcagodme/shortn/protos"
	"google.golang.org/grpc"
)

const (
	port = ":52002"
)

type server struct {
	pb.UnimplementedShortnServer
}

func (s *server) AddPageRpc(ctx context.Context, in *pb.AddPageRequest) (*pb.AddPageResponse, error) {
	log.Printf("AddPageRpc: incoming raw_url = %s", in.GetRawUrl())
	shorturl, err := core.AddPage(in.GetRawUrl())
	if err != nil {
		return &pb.AddPageResponse{ShortUrl: ""}, err
	}
	return &pb.AddPageResponse{ShortUrl: shorturl}, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen to %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterShortnServer(s, &server{})
	log.Println("Started Shortn RPC server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"test-microservice/internal/services"
	"test-microservice/protobuf/short"
)

type ShortenerServer struct {
	short.UnimplementedUrlShortenerServer
	Service services.URLShortener
}

func (s ShortenerServer) ShortenURL(ctx context.Context, req *short.URLRequest) (*short.URLResponse, error) {

	res := &short.URLResponse{URL: "someUrl"}
	return res, nil
}

func (s ShortenerServer) GetOriginalURL(ctx context.Context, req *short.URLRequest) (*short.URLResponse, error) {

	res := &short.URLResponse{URL: "someUrl"}
	return res, nil
}

func gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	short.RegisterUrlShortenerServer(s, &ShortenerServer{
		Service: app.Service,
	})

	log.Printf("gRPC Server started on port %s", gRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

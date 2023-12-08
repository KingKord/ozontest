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
	Shortener services.URLShortener
}

func (s ShortenerServer) ShortenURL(ctx context.Context, req *short.URLRequest) (*short.URLResponse, error) {
	URL := req.GetURL()

	shortenURL, err := s.Shortener.ShortenURL(URL)
	if err != nil {
		return nil, err
	}

	res := &short.URLResponse{URL: shortenURL}
	return res, nil
}

func (s ShortenerServer) GetOriginalURL(ctx context.Context, req *short.URLRequest) (*short.URLResponse, error) {
	shortURL := req.GetURL()

	originalURL, err := s.Shortener.GetOriginalURL(shortURL)
	if err != nil {
		return nil, err
	}

	res := &short.URLResponse{URL: originalURL}
	return res, nil
}

func gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	short.RegisterUrlShortenerServer(s, &ShortenerServer{
		Shortener: app.Service,
	})

	log.Printf("gRPC Server started on port %s", gRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

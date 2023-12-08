package main

import "front-end/protobuf/short"

type ShortenerServer struct {
	short.UnimplementedUrlShortenerServer
}

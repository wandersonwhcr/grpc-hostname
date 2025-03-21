package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"github.com/wandersonwhcr/grpc-hostname/proto"
)

type server struct {
	proto.UnimplementedHostnameServer
}

func main() {
	addr := os.Getenv("HOSTNAME_SERVER_ADDR")

	l, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterHostnameServer(s, &server{})

	log.Printf("Listening: %s", addr)

	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}

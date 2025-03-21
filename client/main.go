package client

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/wandersonwhcr/grpc-hostname/proto"
	"google.golang.org/grpc"
)

func Run() {
	hostname, err := os.Hostname()

	if err != nil {
		log.Fatalf("Failed to Get Hostname: %v", err)
	}

	conn, err := grpc.Dial(os.Getenv("HOSTNAME_SERVER_ADDR"), grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to Connect: %v", err)
	}

	defer conn.Close()

	c := proto.NewHostnameClient(conn)

	ticker := time.NewTicker(time.Second * 2)

	defer ticker.Stop()

	for range ticker.C {
		response, err := c.GetHostname(context.Background(), &proto.GetHostnameRequest{Hostname: hostname})

		if err != nil {
			log.Fatalf("Failed to Get Server Hostname: %v", err)
		}

		log.Printf("Client: %s, Server: %s", hostname, response.Hostname)
	}
}

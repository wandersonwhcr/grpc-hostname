package main

import (
	"log"
	"os"

	"github.com/wandersonwhcr/grpc-hostname/client"
	"github.com/wandersonwhcr/grpc-hostname/server"
)

func main() {
	switch os.Getenv("HOSTNAME_CMD") {
	case "server":
		server.Run()
	case "client":
		client.Run()
	default:
		log.Fatalf("Failed to Command: server or client")
	}
}

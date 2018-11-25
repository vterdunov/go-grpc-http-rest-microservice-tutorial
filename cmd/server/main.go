package main

import (
	"log"

	"github.com/vterdunov/go-grpc-http-rest-microservice-tutorial/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

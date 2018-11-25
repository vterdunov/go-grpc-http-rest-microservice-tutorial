package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"github.com/vterdunov/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
)

func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("Shutting down GRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("Starting GRPC server...")
	return server.Serve(listen)
}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lapitskyss/go_backend_2/server"
)

func main() {
	grpcServer := server.NewGRPCServer()
	grpcServer.Start()

	restServer := server.NewRESTServer()
	restServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		log.Println("Received a signal.", "signal", x.String())
	case err := <-grpcServer.Notify():
		log.Println("Received an error from grpc server.", "err", err)
	case err := <-restServer.Notify():
		log.Println("Received an error from rest server.", "err", err)
	}

	grpcServer.Stop()
	grpcServer.Stop()
}

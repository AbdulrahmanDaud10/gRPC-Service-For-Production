package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var (
	Host = "localhost"
	Port = "5000"
)

func main() {
	addr := fmt.Sprintf("%s:%s", Host, Port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	log.Println("tcp listener started at port: ", Port)
	grpcServer := grpc.NewServer()

	// TODO: Implementing and Registering Geometry into gRPC server

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}

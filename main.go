package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/amalaruja/go-hello-grpc/hello"
	"github.com/amalaruja/go-hello-grpc/server"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server.HelloServer{Message: "DC 2"}

	grpcServer := grpc.NewServer()
	hello.RegisterHelloGrpcServiceServer(grpcServer, &s)
	grpcServer.Serve(lis)
}

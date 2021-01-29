package server

import (
	"fmt"

	"golang.org/x/net/context"
	"github.com/amalaruja/go-hello-grpc/hello"
)

type HelloServer struct {
	Message string
}

func (pls *HelloServer) Hello (ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	replyMessage := fmt.Sprintf("Hello from %s", pls.Message)
	return &hello.HelloResponse{Reply: replyMessage}, nil
}
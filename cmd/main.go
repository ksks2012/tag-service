package main

import (
	"log"
	"net"

	pb "github.com/tag-service/proto"
	"github.com/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	port = "18080"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}

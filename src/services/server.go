package main

import (
	"XrBinary/src/xrbinarysrvpb"
	"context"
	"fmt"
	"log"
	"net"
	"xrbinarypb"
	"xrbinarysrvpb"

	"github.com/docker/docker/api/server/router/grpc"
	"google.golang.org/grpc"
)
type server struct {
	xrbinarysrvpb.UnimplementedUploadServiceServer
}

func (s *server) UploadRequest( ctx context.Context , req *xrbinarysrvpb.XrDebugRequest) (*xrbinarysrvpb.XrDebugResponse, error){
	fmt.Println("something")
	return &xrbinarysrvpb.XrDebugResponse{Result: true,Error: "None"}, nil
}

func main () {
	server_port := "50051"
	server_ip := "0.0.0.0"
	fmt.Printf("starting gRPC server %s at %s " , server_ip, server_port)
	listner, err := net.Listen("tcp", server_ip+ ":" +server_port)
	if err != nil {
		log.Fatalf("failed to start server on : %v", err)
	}
	s := grpc.NewServer()
	xrbinarysrvpb.RegisterUploadServiceServer(s, &server{} )
	log.Printf("server listening at %v", listner.Addr())
	if err := s.Serve(listner); err != nil {
		log.Fatalf("Failer to server GRPC %v", err)
	}
}
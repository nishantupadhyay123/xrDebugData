package main

import (
	pb "github.com/nishantupadhyay123/xrDebugData/src/xrbinarysrvpb"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"io"
)
type server struct {
	pb.UnimplementedUploadServiceServer
}


func (*server) UploadRequest(stream pb.UploadService_UploadRequestServer) error {
	fmt.Printf("UploadRequest function was invoked with a streaming request\n")
	result := false
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&pb.XrDebugResponse{
				Result: result,
				Error: "none",
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		request_id := req.GetReqid()
		fmt.Println(" The requestid is %v", request_id)
		result = true
	}
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
	pb.RegisterUploadServiceServer(s, &server{})
	log.Printf("server listening at %v", listner.Addr())
	if err := s.Serve(listner); err != nil {
		log.Fatalf("Failer to server GRPC %v", err)
	}
}
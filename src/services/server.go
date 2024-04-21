package main

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
	datapb "github.com/nishantupadhyay123/xrDebugData/src/xrbinarypb"
	pb "github.com/nishantupadhyay123/xrDebugData/src/xrbinarysrvpb"
	"google.golang.org/grpc"
)

const parentDir = "/Users/niupadhy/Downloads"
type server struct {
	pb.UnimplementedUploadServiceServer
}


func (*server) UploadRequest(stream pb.UploadService_UploadRequestServer) error {
	log.Printf("UploadRequest function was invoked with a streaming request\n")
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
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		request_id := req.GetReqid()
		log.Printf(" The requestid is %v\n", request_id)
		log.Print("passing to the decoder\n")
		error := DecodeAndSave(req.Data)
		if error != nil {
			result = true
		}
	}
}
func DecodeAndSave (pbdata []byte) (error) {
	//may be we should process buffer size
	// var namebuf bytes.Buffer
	xrdata := &datapb.XrDebugData{}
	err := proto.Unmarshal(pbdata , xrdata)
	if err!=nil {
		log.Fatalf("Cant decode the data as per protobuf %v\n",err)
	}
	data_dir :=  xrdata.GetStreamId() + "_" + xrdata.GetHostName() + "_" + xrdata.MsgType.String()
	node_dir := xrdata.GetLtrace().GetNodeName()
	path := filepath.Join(parentDir,data_dir, node_dir)
	log.Printf("path is %v",path)
	err = os.MkdirAll(path , 0755)
	os.Chdir(path)
	f, err := os.OpenFile("ofa", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )
	if err!= nil{
		log.Fatalf("Not ale to open the file %v\n", err)
	}
	for _,trace := range xrdata.Ltrace.GetLtrace(){
		_,err := f.Write(trace.GetMsgData())
		if err!=nil{
			log.Fatalf("Couldnt write the file :%v\n", err)
		}
	}
	err = f.Close()
	if err!= nil {
		log.Printf("couldnt close the file\n")
	}
	return nil

}

func main () {
	server_port := "50051"
	server_ip := "0.0.0.0"
	log.Printf("starting gRPC server %s at %s \n" , server_ip, server_port)
	listner, err := net.Listen("tcp", server_ip+ ":" +server_port)
	if err != nil {
		log.Fatalf("failed to start server on : %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterUploadServiceServer(s, &server{})
	log.Printf("server listening at %v\n", listner.Addr())
	if err := s.Serve(listner); err != nil {
		log.Fatalf("Failer to server GRPC %v\n", err)
	}
}
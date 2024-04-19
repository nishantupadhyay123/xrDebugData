package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/nishantupadhyay123/xrDebugData/src/xrbinarysrvpb"
	datapb "github.com/nishantupadhyay123/xrDebugData/src/xrbinarypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// type File struct {
// 	Files []string
// 	CharacterEncoding string

// 	filenames []string
// 	decoder *encoding.TextMarshaler
// }

//Red the bytes in 4k size only for efficiency
const BufferSize = 500


func doStream(c pb.UploadServiceClient){
	log.Printf("starting the client\n")
	requests := []*pb.XrDebugRequest{
		&pb.XrDebugRequest{
		Reqid: 1,
		Data: []byte("I am nishant"),
		Errors: "None",
		Decoder: "None",
	},
	&pb.XrDebugRequest{
		Reqid: 2,
		Data: []byte("I am Rajni"),
		Errors: "None",
		Decoder: "None",
	},
}
stream, err := c.UploadRequest(context.Background())
for _,req := range requests{
	log.Printf("Sending req: %v\n", req.Reqid)
	stream.Send(req)
	time.Sleep(1000 * time.Millisecond)
}
res, err := stream.CloseAndRecv()
if err!=nil{
	log.Fatalf("Error in sending \n ")
}
log.Printf("stream response %v", res)
}

func  doStreamFile(c pb.UploadServiceClient){
	log.Printf("Sending messahge")

	filepath := "/Users/niupadhy/Documents/Study/Prep/godownload"
	files, err := ReadDirectory(filepath)
	if err != nil {
		log.Fatalf("No files found or directory %v\n", err)
	}
	log.Print(files)
	for _,file := range files {
		log.Printf("Sending file %v\n", file)
		file, err := os.Open(file)
		if err != nil {
			log.Print("failed to open file %v : %v\n", file, err)
			continue
		}
		defer file.Close()

		//May be needed to create  a parallel send.
		// filestat, err := file.Stat()
		filesize, err := file.Stat()
		if err != nil {
			log.Print("not able to get data bout file")
		}
		//for information purposes.
		log.Printf("Size of the %v :%v",file.Name(), filesize.Size())

		//make a slice for buffer read.
		requests := []*pb.XrDebugRequest{}
		i := 0
		for {
			buffer := make([]byte,BufferSize)
			bytesread, err := file.Read(buffer)
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				}
				break
			}
			i++;
			x := pb.XrDebugRequest{
				Reqid: int64(i),
				Data: buffer[:bytesread],
				Errors: "None",
				Decoder: file.Name(),
			}
			requests = append(requests, &x)
		}
		stream, err := c.UploadRequest(context.Background())
		for _,req := range requests{
			log.Printf("Sending req: %v\n", req.Reqid)
			stream.Send(req)
			time.Sleep(10 * time.Millisecond)
		}
		res, err := stream.CloseAndRecv()
		if err!=nil{
			log.Fatalf("Error in sending \n ")
		}
		log.Printf("stream response %v", res)
	}
}


func ReadDirectory (dir string) ([]string , error) {
	// Doesn't support file walk in sub-directory structure.
	// Should use filepath.walk.
	var files []string
	f, err := os.Open(dir)
	if err != nil {
		log.Fatalf("The directory doesnt exist or not accessible %v", err)
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        return files, err
    }
    for _, file := range fileInfo {
		filestring := dir + "/" + file.Name()
        files = append(files, filestring)
    }
    return files, nil
}

func main() {
	fmt.Println("Hello I'm a xrDebugClient client")
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	// tls := false
	// if tls {
	// 	certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
	// 	creds,err := insecure.NewCredentials().
	// 	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	// 	if sslErr != nil {
	// 		log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
	// 		return
	// 	}
	// 	opts = grpc.WithTransportCredentials(creds)
	// }

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewUploadServiceClient(cc)
	fmt.Printf("Created client")

	// Simple message
	doStream(c)
	// Send file
	doStreamFile(c) 

	// sent thru protobuf a text file.
	// doStreamProtoBufPlainText(c)

	// //stringZippedFile in protbuf
	// doStreamProtoBufPlainZip(c)
}
package main

import (
	// "bytes"
	"context"
	"net"
	// "encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	datapb "github.com/nishantupadhyay123/xrDebugData/src/xrbinarypb"
	pb "github.com/nishantupadhyay123/xrDebugData/src/xrbinarysrvpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

// type File struct {
// 	Files []string
// 	CharacterEncoding string

// 	filenames []string
// 	decoder *encoding.TextMarshaler
// }

//Red the bytes in 4k size only for efficiency
const BufferSize = 500

//this is sample code.
const LC = "RP0"



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
	log.Printf("Sending message")

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
			time.Sleep(5 * time.Millisecond)
		}
		res, err := stream.CloseAndRecv()
		if err!=nil{
			log.Fatalf("Error in sending \n ")
		}
		log.Printf("stream response %v", res)
	}
}


func  doStreamProtoBufPlainText(c pb.UploadServiceClient){
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

		/*
		Create the Protobuff data 
		*/
		data := &datapb.XrDebugData{}
		host, err := os.Hostname()
		if err != nil {
			log.Print("couldnt identify the hostname \n")
		}
		data.HostId = &datapb.XrDebugData_HostName{
			HostName: host,
		}
		data.StreamId = "testcase1"
		data.MsgType = datapb.XrDebugData_LTRACE
		data.Ltrace = &datapb.XrLtrace{}
		data.Ltrace.Node = &datapb.XrLtrace_NodeName{
			NodeName: LC,
		}
		process_data := []*datapb.ProcessTrace{}
		for {
			//control from client side to send the buffer length
			// need to build server side cle
			buffer := make([]byte,BufferSize)
			bytesread, err := file.Read(buffer)
			if err!= nil {
				if err != io.EOF {
					log.Println(err)
				}
				break
			}
			x := datapb.ProcessTrace{
				EventType: "Something",
				Data: &datapb.ProcessTrace_MsgData{
					MsgData : buffer[:bytesread],
				},
			}
			process_data = append(process_data, &x)
		}
		data.Ltrace.Ltrace = process_data
		/*
		end of the Protobuf send.
		*/
		// log.Printf("value is %v", data)
		byteData,err := proto.Marshal(data)
		if err !=nil{
			log.Fatalf("Couldnt marshal the data %v", err)
		}

		// var to_Send bytes.Buffer
		// enc := gob.NewEncoder(&to_Send)
		// err := enc.Encode(*data)


		// /*create the server side proto*/
		y := &pb.XrDebugRequest{
			Reqid: 1,
			Data: byteData,
			Errors: "None",
			Decoder: "Ltrace Data",
		}

		stream,err := c.UploadRequest(context.Background())
		stream.Send(y)
		res, err := stream.CloseAndRecv()
		if err!=nil{
			log.Fatalf("Error in sending \n ")
		}
		log.Printf("Returned status is %v",res)
}
}

func ReadDirectory(dir string) ([]string , error) {
	// Doesn't support file walk in sub-directory structure.
	// Should use filepath.walk.
	var files []string
	f, err := os.Open(dir)
	if err != nil {
		log.Fatalf("The directory doesnt exist or not accessible %v\n", err)
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

func handle_data (conn net.Conn , c pb.UploadServiceClient){
	fmt.Println("connected to : ", conn.RemoteAddr())
	/*
	Create the Protobuff data 
	*/
	data := &datapb.XrDebugData{}
	host, err := os.Hostname()
	if err != nil {
		log.Print("couldnt identify the hostname \n")
	}
	data.HostId = &datapb.XrDebugData_HostName{
		HostName: host,
	}
	data.StreamId = "testcase2"
	data.MsgType = datapb.XrDebugData_LTRACE
	data.Ltrace = &datapb.XrLtrace{}
	data.Ltrace.Node = &datapb.XrLtrace_NodeName{
		NodeName: LC,
	}
	process_data := []*datapb.ProcessTrace{}
	for {
		//control from client side to send the buffer length
		// need to build server side cle
		buffer := make([]byte,BufferSize)
		bytesread, err := conn.Read(buffer)
		if err!= nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		x := datapb.ProcessTrace{
			EventType: "Something",
			Data: &datapb.ProcessTrace_MsgData{
				MsgData : buffer[:bytesread],
			},
		}
		process_data = append(process_data, &x)
	}
	data.Ltrace.Ltrace = process_data
	/*
	end of the Protobuf send.
	*/
	// log.Printf("value is %v", data)
	byteData,err := proto.Marshal(data)
	if err !=nil{
		log.Fatalf("Couldnt marshal the data %v", err)
	}

	// var to_Send bytes.Buffer
	// enc := gob.NewEncoder(&to_Send)
	// err := enc.Encode(*data)


	// /*create the server side proto*/
	y := &pb.XrDebugRequest{
		Reqid: 1,
		Data: byteData,
		Errors: "None",
		Decoder: "Ltrace Data",
	}

	stream,err := c.UploadRequest(context.Background())
	stream.Send(y)
	res, err := stream.CloseAndRecv()
	if err!=nil{
		log.Fatalf("Error in sending \n ")
	}
	log.Printf("Returned status is %v",res)
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
	fmt.Printf("Created  grpc client")
	//doStreamProtoBufPlainText(c)
	

	server_port := "50052"
	server_ip := "0.0.0.0"
	log.Printf("starting simple socket listner at %s at %s \n" , server_ip, server_port)
	listner, err := net.Listen("tcp", server_ip+ ":" +server_port)
	if err != nil {
		log.Fatalf("failed to start listner on : %v\n", err)
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalln("Error while accepting")
		}
		go handle_data(conn, c)
	}


	// Simple message
	// doStream(c)
	// // // Send file
	// doStreamFile(c) 

	// sent thru protobuf a text file.
	//doStreamProtoBufPlainText(c)

	// //stringZippedFile in protbuf
	// doStreamProtoBufPlainZip(c)
}
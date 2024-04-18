package main
import (
	pb "github.com/nishantupadhyay123/xrDebugData/src/xrbinarysrvpb"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"io"
)

func main() {
	fmt.Println("Hello I'm a xrDebugClient client")

	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	// cc, err := grpc.Dial("localhost:50051", opts)
	cc, err := grpc.D

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)

}

func strea
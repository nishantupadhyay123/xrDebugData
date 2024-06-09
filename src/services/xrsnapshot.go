/* ----------------------------------------------------------------------------
* xr_debug_client - XR debug data TCP client
* Author : Nishant Upadhyay
* Arpil 2024,Copyright (c) 2016 by Cisco Systems, Inc.
* ----------------------------------------------------------------------------
 */
package XrDebugData

import (
	"github.com/nishantupadhyay123/xrDebugData/src/xrbinarysrvpb"
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"log"
	"net"
	"sync"
	"text/scanner"

	"github.com/golang/protobuf/protoc-gen-go/grpc"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"google.golang.org/grpc"
)

type  XrDebugData struct {
	// common configuration
	Transport      string 
	serviceAddress string `toml:"service_address"`
	listner        net.Listener

	// TCP related configuration
	AllowedPendingMessages int `toml:"allowed_pending_messages"`
	MaxTCPConnections      int `toml:"max_tcp_connections"`
	

	// Internal state 
	all_conn     []*net.Conn
	acc   		   telegraf.Accumulator
	cancel 		   context.CancelFunc
	ctx 		   context.Context
	wg    		   sync.WaitGroup			
}

type server struct{
	*xrbinarysrvpb.UnimplementedUploadServiceServer
}

// refrenced to https://pkg.go.dev/github.com/influxdata/telegraf#ServiceInput 

func (x *XrDebugData) Start (acc telegraf.Accumulator) error {
	var err error
	x.acc = acc
	x.ctx, x.cancel = context.WithCancel(context.Background())

	switch x.Transport {
	case "TCP":
		x.listner, err = net.Listen("tcp", x.serviceAddress)
		if err != nil {
			return err
		}
		x.wg.Add(1)
		go x.acceptTCPClients()
		log.Printf(" Start the xrDebug Plugin at : %s", x.serviceAddress)

	case "GRPC":
		var opts []grpc.ServerOption
		// Initialize grpC server.
		x.listner, err = net.Listen("tcp", x.serviceAddress)
		if err != nil {
			return err
		}
		//no TLS supported as of now but just creating blank
		x.wg.Add(1)
		go func() {
			grpcserver := grpc.NewServer(opts...)
			xrbinarysrvpb.RegisterUploadServiceServer(grpcserver, &server{})	
		}()
	}
	
	return nil
}

func (x *XrDebugData) acceptTCPClients() error {
	// Add the code to accept tcp
	// send the payload buffer to decode and store
	var mutex sync.Mutex
	clients := make(map[net.Conn]struct{})
	for {
		conn , err := x.listner.Accept()
		if err != nil {
			if x.ctx.Err() != nil {
				break
			}
			x.acc.AddError(fmt.Errorf("E! Failed to accept TCP connection: %v", err))
		}
		mutex.Lock()
		clients[conn] = struct{}{}
		mutex.Unlock()
		// Process 1 connection 
		x.wg.Add(1)
		go func (){
			log.Printf("D! Accepted Cisco XrDebugData connection from %s", conn.RemoteAddr())
			scanner := bufio.NewScanner(conn)
			x.decodeMessage(payload.Bytes())

		}()
		log.Printf("D! Closed TCP dialout connection from %s", conn.RemoteAddr())
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
		x.wg.Done()
	}
	return nil
}


// decode the bytes according to protobuf.
func (x *XrDebugData) decodeMessage(data []byte) error {
	return nil
}

// ##### telegraf input plugin settings

const sampleConfig = `
 ## transport type (oneof TCP and future gRPC)
 Transport = "TCP"

 ## Port on which XrDebugData will listen to
 service_address : "5900" 

 ## No. of tcp outstanding messages be processed.
 allowed_pending_messages : 100000
  
 ## Max. no of TCP connections supported.Need to do a little research but we should tune this.
 max_tcp_connections = 200
`

// SampleConfig of plugin
func (x *XrDebugData) SampleConfig() string {
	return sampleConfig
}

// Description of plugin
func (x *XrDebugData) Description() string {
	return "IOS-XR debug data plugin"
}

func ( x *XrDebugData) Gather (_ telegraf.Accumulator) error {
	return nil
}


func init() {
	inputs.Add("xrdebugdata", func() telegraf.Input {
		return &XrDebugData{
			Transport:      "TCP", // future could be gRPC.
			serviceAddress: ":5900",
			AllowedPendingMessages: 100000,
			MaxTCPConnections: 200,
		}
	})
}


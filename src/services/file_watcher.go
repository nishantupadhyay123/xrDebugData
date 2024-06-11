/* ----------------------------------------------------------------------------
* Watches the file and trasnfer to  RP to transmit outside.
* Author : Nishant Upadhyay
* Arpil 2024,Copyright (c) 2016 by Cisco Systems, Inc.
* ----------------------------------------------------------------------------
 */
package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)
type fileData struct {
	filesize int64
	transferred bool
}

type localData struct {
	transferred_list  map[string]fileData
	connList []net.Conn
	LC string
}

var Mutex sync.Mutex

const BufferRead = 500

func CreateConn() (data localData , err error){
	//initialize the local data
	p := data
	p.transferred_list = make(map[string]fileData)

	//create the 3 connections.
	i := 0
	for i < 3 {
		var Dial net.Dialer
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		var conn net.Conn
		var err error
		for {
			conn, err = Dial.DialContext(ctx, "tcp", "localhost:50052")
			if err != nil {
				log.Printf("Failed to dial: %v", err)
				// Below lines should be replaced by backoff algorithm
				log.Print("trying again after 5 seconds")
				d := 5 * time.Second
				time.Sleep(d)
				continue
			}
			break;
		}
		i++
		p.connList = append(p.connList, conn)
	}
	return p, nil
}

//send when needed . There is no need for to keep this open
func fileSend(filename fsnotify.Event, data localData) {
	Mutex.Lock()
	defer Mutex.Unlock()
	connst := &data
	conns := connst.connList[rand.Intn(len(connst.connList))]
	log.Printf("selected connection is %v",conns)
	log.Println(conns.LocalAddr())
	// defer conns.Close() 
	file , err := os.Open(filename.Name)
	defer file.Close()
	filesize, err := file.Stat()
	if err != nil {
		log.Println("not able to get data about the file %s", file.Name())
	}
	y := fileData{}
	y.filesize = filesize.Size()
	for {
		buffer := make([]byte, BufferRead)
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		_, err = conns.Write(buffer[:bytesRead])
		if err != nil{
			log.Fatalln("Error while writing to RP %v",err)
		}
	}
	y.transferred = true
	connst.transferred_list[file.Name()] = y
}

func main() {
	

    //
	p, err := CreateConn()
	if err != nil {
		log.Fatal("couldnt be established")
	}
	log.Println("total created %v",p)
    // Create new watcher.
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    // Start listening for events.
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                log.Println("event:", event)
                if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
                    log.Println("modified file:", event.Name)
					log.Printf("Transferring the file")
					p.LC = "0/RP0/CPU0"
					go fileSend(event, p)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    // Add a path.
    err = watcher.Add("/Users/niupadhy/Downloads/junk")
    if err != nil {
        log.Fatal(err)
    }

    // Block main goroutine forever.
    <-make(chan struct{})
}
package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// Server RPC

type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, replay *int64) error {
	// Fill reply pointer to send the data back
	*replay = time.Now().Unix()
	return nil
}

func main() {
	timeServer := new(TimeServer)

	_ = rpc.Register(timeServer)
	rpc.HandleHTTP()

	// Listen for request on port 1234
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	_ = http.Serve(l, nil)
}

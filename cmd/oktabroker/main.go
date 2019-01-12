package main

import (
	"fmt"
	"log"
//	"errors"
	"net"
	
	"github.com/identitybroker/api"
	"google.golang.org/grpc"

)

// Starts gRPC server and waits for connection
func main() {
	
	// Create socket
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := api.Server{}
	
	grpcServer := grpc.NewServer()
	
	//regsiter the ping server
	api.RegisterPingServer(grpcServer, &s)
	
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}
}


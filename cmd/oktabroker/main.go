package main

import (
	"fmt"
	"log"
//	"errors"
	"net"
	
	api "github.com/identitybroker/api/handler"
	provider "github.com/identitybroker/pkg/provider"
	"google.golang.org/grpc"

)

// Starts gRPC server and waits for connection
func main() {
	
	// Create socket
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := provider.ResourceProviderService{}
	grpcServer := grpc.NewServer()
	
	//regsiter the resource provider server
	api.RegisterResourceProviderServer(grpcServer, s)
	
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}
}
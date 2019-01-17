package main

import (
	"fmt"
	"log"
//	"errors"
	"net"
	
	rpcgen "github.com/identitybroker/api/_generated"
	api "github.com/identitybroker/api/handler"
	mapper "github.com/identitybroker/api/mapper"
	"google.golang.org/grpc"

)

// Starts gRPC server and waits for connection
func main() {
	
	// Create socket
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := api.ResourceProviderService{ResMapper: mapper.NewResourceMapper()}
	grpcServer := grpc.NewServer()
	
	//regsiter the resource provider server
	rpcgen.RegisterResourceProviderServer(grpcServer, s)
	
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}
}
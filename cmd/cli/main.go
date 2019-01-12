package main

import (
	"log"
//	"errors"
	"golang.org/x/net/context"
	
	"github.com/identitybroker/api"
	"google.golang.org/grpc"

)

// Starts gRPC server and waits for connection
func main() {
	var conn *grpc.ClientConn
	
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connot reach the server. Failed with erro : %v", err)
	}
	defer conn.Close()
	
	c := api.NewPingClient(conn)
	
	// Invoke API
	r, err := c.SayHello(context.Background(), &api.PingMessage{Greetings: "Hello World!"})
	if err != nil {
		log.Fatalf("Error invoking API : %v", err)
	}	
	
	log.Printf("Response from the server : %s\n", r.Greetings)
}


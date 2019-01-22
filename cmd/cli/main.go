package main

import (
	"log"
//	"errors"
	"golang.org/x/net/context"
	
	rpcgen "github.com/identitybroker/api/_generated"
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
	
	c := rpcgen.NewResourceProviderClient(conn)
	
	// Invoke API
	opt := rpcgen.ProviderRegistrationOpts{
		Namespace: "myOkta1", 
		ProviderType: "Okta", 
		Cred: &rpcgen.Credential{
			ApiKey: "goodkey",
			AuthUrl: "goodUrl"},
	}
	r, err := c.RegisterProvider(context.Background(), &opt)
	if err != nil {
		log.Fatalf("Error invoking API : %v", err)
	}	
	
	log.Printf("Response from the server : %s\n", r.ProviderId)
}


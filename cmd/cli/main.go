package main

import (
	"log"
	"fmt"
	"time"
//	"errors"
    "math/rand"
	"golang.org/x/net/context"
	
	"github.com/identitybroker/api/_generated"
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
	
	c := api.NewResourceProviderClient(conn)
	
	// Invoke API
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)	
	i := r1.Intn(100)
	opt := api.ProviderRegistrationOpts{
				Namespace: fmt.Sprintf("Okta-dev%d.preview.com", i), 
				ProviderType: "Okta", 
				Cred: &api.Credential{
						ApiKey: "d2eadd4b-9528-4adb-bcce-018956c39a5a",
						AuthUrl: fmt.Sprintf("https://Okta-dev%d.preview.com/%s", i, "authorize")},
		}
	r, err := c.RegisterProvider(context.Background(), &opt)
	if err != nil {
		log.Fatalf("Error invoking API : %v", err)
	}	
	
	// Print Status
	log.Println(r)
	
	log.Printf("Trying to update provider : %s(%s)\n", r.ProviderNamespace, r.ProviderId)
							
	// Updating provider token
	updateOpt := &api.ProviderUpdateOpts{
					ProviderId	: r.ProviderId,
					ProviderName: r.ProviderNamespace,
					Cred		: &api.Credential{ApiKey: "02e25f8a-827d-44c8-adfc-c086ffa878e6",
									 AuthUrl: "https://Okta-dev100.preview.com/authorize",
								},
	}
	s, _ := c.UpdateProvider(context.Background(), updateOpt)
	
	// Print Status
	log.Println(s)
}


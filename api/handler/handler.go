package api

import (
	"log"
	
	"golang.org/x/net/context"
)

// gRPC Server handler
type Server struct {}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive Message %s\n", in.Greetings)
	return &PingMessage{Greetings: in.Greetings}, nil
}
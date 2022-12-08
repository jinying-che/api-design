package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/jinying-che/api-design/grpc/go/ping"
)

// define the service
type service struct{}

// implement the SayHello RPC method
func (s *service) SayHello(ctx context.Context, in *ping.HelloRequest) (*ping.HelloResponse, error) {
	ret := "Hello " + in.Name
	log.Printf("server SayHello: %s", ret)
	return &ping.HelloResponse{Message: ret}, nil
}

// implement the GetAge RPC method
func (s *service) GetAge(ctx context.Context, in *ping.AgeRequest) (*wrappers.Int32Value, error) {
	log.Printf("server Age: %d", in.Age)
	return &wrappers.Int32Value{Value: in.Age}, nil
}

// implement the CheckHealth RPC method
func (s *service) CheckHealth(ctx context.Context, in *empty.Empty) (*ping.HealthResponse, error) {
	log.Printf("server Health: %s", "OK")
	return &ping.HealthResponse{Status: "OK"}, nil
}

func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	s := grpc.NewServer()

	// register the service with the server
	ping.RegisterPingServer(s, &service{})

	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

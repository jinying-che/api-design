package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinying-che/api-design/grpc/go/ping"
)

const (
	address     = "localhost:7777"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := ping.NewPingClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &ping.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	age := int32(25)

	if len(os.Args) > 2 {
		inputAge, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("invalid age: %v", err)
		}
		age = int32(inputAge)
	}

	ageReq := &ping.AgeRequest{Age: age}
	ageResp, err := c.GetAge(ctx, ageReq)

	if err != nil {
		log.Fatalf("could not get age: %v", err)
	}

	log.Printf("Age: %d", ageResp.Value)

	empty := &empty.Empty{}
	healthResp, err := c.CheckHealth(ctx, empty)

	if err != nil {
		log.Fatalf("could not check health: %v", err)
	}

	log.Printf("Health: %s", healthResp.Status)
}

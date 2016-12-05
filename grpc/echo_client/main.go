package main

import (
	"log"
	"os"

	pb "github.com/ttallskog/sample-golang/grpc/messages"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address        = "localhost:50051"
	defaultMessage = "I am runnig"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	// Contact the server and print out its response.
	message := defaultMessage
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	r, err := c.Echo(context.Background(), &pb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Echo: %s", r.Message)
}

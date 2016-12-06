package main

import (
	"flag"
	"log"
	"strconv"

	"time"

	pb "github.com/ttallskog/sample-golang/grpc/messages"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const (
	address        = "localhost:50051"
	defaultMessage = "I am call: "
)

// LoopEcho maskes x number of sequential calls via grpc
func LoopEcho(client pb.EchoClient, loops int) error {
	for i := 0; i < loops; i++ {
		r, err := client.Echo(context.Background(), &pb.EchoRequest{Message: defaultMessage + strconv.Itoa(i+1)})
		if err != nil {
			return err
		}
		log.Printf("Echo: %s", r.Message)
	}
	return nil
}

// Timer is a time logging function
func Timer(message string) func() {
	s := time.Now()
	log.Println(message, "started")
	return func() {
		e := time.Now().Sub(s)
		log.Println(message, "took", e)
	}
}

func main() {
	stop := Timer("grpc call(s)")
	defer stop()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	numbPtr := flag.Int("calls", 1, "an int")
	flag.Parse()

	e := LoopEcho(pb.NewEchoClient(conn), *numbPtr)
	if e != nil {
		log.Fatalf("could not greet: %v", e)
	}
	log.Printf("complete")
}

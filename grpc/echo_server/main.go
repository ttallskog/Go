package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/ttallskog/sample-golang/grpc/messages"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	return &pb.EchoReply{Message: "Content: " + in.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

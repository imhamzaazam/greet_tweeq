package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"greet/server/greet/pb"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGreetServer
}

func (s *server) SayGreet(ctx context.Context, empty *emptypb.Empty /*in *pb.SayGreetRequest*/) (*pb.SayGreetResponse, error) {
	log.Printf("Received: %v" /*in.GetName()*/)
	return &pb.SayGreetResponse{
		Message: "Hello I am the wonderful engineer who joined Tweeq",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterGreetServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

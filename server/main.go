package main

import (
	"context"
	"log"
	"net"

	"github.com/aymene01/my_grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.A + in.B,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln("Fail to run server", err)
	}

	s := grpc.NewServer()
	reflection.Register((s))

	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Fail to serve", err)
	}
}

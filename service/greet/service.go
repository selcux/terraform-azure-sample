package greet

import (
	"context"
	"fmt"
	"log"
	"net"

	greetgrpc "github.com/selcux/terraform-azure-sample/pkg/grpc/greet"
	"google.golang.org/grpc"
)

type Service struct {
	gs *grpc.Server
	greetgrpc.UnimplementedGreetServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Run(host string, port int) error {
	s.gs = grpc.NewServer()
	greetgrpc.RegisterGreetServiceServer(s.gs, s)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("unable to listen %v", err)
	}

	log.Println("Starting GRPC server...")

	return s.gs.Serve(lis)
}

func (s *Service) SayHello(_ context.Context, request *greetgrpc.GreetRequest) (*greetgrpc.GreetResponse, error) {
	return &greetgrpc.GreetResponse{
		Message: fmt.Sprintf("Hello %s!!", request.Name),
	}, nil
}

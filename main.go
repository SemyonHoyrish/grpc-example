package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/SemyonHoyrish/grpc-example/protos/routes/user"
	"google.golang.org/grpc"
)

type userService struct {
	pb.UnimplementedUserServer
}

func (s *userService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		ID:     req.ID,
		Name:   "Test",
		Email:  "test@test",
		Status: "i am good",
		Online: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &userService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

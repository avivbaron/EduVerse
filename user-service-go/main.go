package main

import (
	"context"
	"log"
	"net"

	pb "github.com/avivbaron/eduverse/proto/userpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Id:    req.id,
		Name:  "John Doe",
		Email: "john@mail.com",
	}, nil
}

func (s *server) ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.ValidationResponse, error) {
	return &pb.ValidationResponse{IsValid: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("gRPC UserService running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

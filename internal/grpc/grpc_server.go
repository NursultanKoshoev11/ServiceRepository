package grpc

import (
	"context"
	"log"
	"net"
	"servicerepository/internal/service"

	grpcrepo "github.com/NursultanKoshoev11/GeneralProtoContracts/generated/repository"
	"google.golang.org/grpc"
)

type gRPCService struct {
	grpcrepo.UnimplementedUserServiceServer
	service *service.UserService
}

func (s *gRPCService) CreateUser(ctx context.Context, req *grpcrepo.CreateUserRequest) (*grpcrepo.UserResponse, error) {

	user, err := s.service.Register(req.Name, req.Email, req.Password)

	if err != nil {
		log.Println("cant to register ", err)
		return nil, err
	}

	return &grpcrepo.UserResponse{
		User: &grpcrepo.User{
			Id:       int32(user.ID),
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil
}

func RunGRPCServer(svc *service.UserService) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpcrepo.RegisterUserServiceServer(grpcServer, &gRPCService{
		service: svc,
	})

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

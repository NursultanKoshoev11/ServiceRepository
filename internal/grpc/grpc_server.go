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
	grpcrepo.UnimplementedRepositoryServiceServer
	userservice *service.UserService
}

func (grpcservice *gRPCService) GetUserByEmail(ctx context.Context, req *grpcrepo.GetUserByEmailRequest) (*grpcrepo.UserResponse, error) {
	user, err := grpcservice.userservice.GetUserByEmail(req.Email)

	if err != nil {
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

func (grpcservice *gRPCService) GetUserByID(ctx context.Context, req *grpcrepo.GetUserByIDRequest) (*grpcrepo.UserResponse, error) {

	user, err := grpcservice.userservice.GetUserByID(int64(req.Id))

	if err != nil {
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

func (grpcservice *gRPCService) GetUserByName(ctx context.Context, req *grpcrepo.GetUserByNameRequest) (*grpcrepo.UserResponse, error) {

	user, err := grpcservice.userservice.GetUserByName(req.Name)

	if err != nil {
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

func (grpcservice *gRPCService) CreateUser(ctx context.Context, req *grpcrepo.CreateUserRequest) (*grpcrepo.UserResponse, error) {

	user, err := grpcservice.userservice.Register(req.Name, req.Email, req.Password)

	if err != nil {
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
	grpcrepo.RegisterRepositoryServiceServer(grpcServer, &gRPCService{
		userservice: svc,
	})

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

package grpc

import (
	"context"
	"log"
	"net"
	"servicerepository/config"
	"servicerepository/internal/models"
	"servicerepository/internal/service"

	grpcrepo "github.com/NursultanKoshoev11/GeneralProtoContracts/generated/repository"
	"google.golang.org/grpc"
)

type gRPCService struct {
	grpcrepo.UnimplementedRepositoryServiceServer
	userservice *service.UserService
}

func (grpcservice *gRPCService) CreateRole(ctx context.Context, req *grpcrepo.CreateRoleRequest) (*grpcrepo.ErrorResponse, error) {
	return nil, grpcservice.userservice.CreateRole(models.RoleType(req.RoleID), req.RoleName)
}
func (grpcservice *gRPCService) CreateUser(ctx context.Context, req *grpcrepo.CreateUserRequest) (*grpcrepo.User, error) {
	user, err := grpcservice.userservice.CreateUser(req.Email, req.Email, models.RoleType(req.RoleID))

	if err != nil {
		return nil, err
	}

	return &grpcrepo.User{
		Id:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		RoleID:    int32(user.RoleID),
		CreatedAt: user.CreatedAt.String(),
	}, nil
}
func (grpcservice *gRPCService) DeleteProfileByUserID(ctx context.Context, req *grpcrepo.GetUserByIDRequest) (*grpcrepo.ErrorResponse, error) {
	return nil, grpcservice.userservice.DeleteProfileByUserID(req.Id)
}
func (grpcservice *gRPCService) DeleteUserByEmail(ctx context.Context, req *grpcrepo.GetUserByEmailRequest) (*grpcrepo.ErrorResponse, error) {
	return nil, grpcservice.userservice.DeleteUserByEmail(req.Email)
}
func (grpcservice *gRPCService) DeleteUserByID(ctx context.Context, req *grpcrepo.GetUserByIDRequest) (*grpcrepo.ErrorResponse, error) {
	return nil, grpcservice.userservice.DeleteUserByID(req.Id)
}
func (grpcservice *gRPCService) GetProfileByEmail(ctx context.Context, req *grpcrepo.GetUserByEmailRequest) (*grpcrepo.Profile, error) {
	profile, err := grpcservice.userservice.GetProfileByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	return &grpcrepo.Profile{
		Id:        profile.ID,
		UserID:    profile.UserID,
		Name:      profile.Name,
		Avatar:    profile.Avatar,
		Bio:       profile.Bio,
		CreatedAt: profile.CreatedAt.String(),
	}, nil
}
func (grpcservice *gRPCService) GetProfileByUserID(ctx context.Context, req *grpcrepo.GetUserByIDRequest) (*grpcrepo.Profile, error) {
	profile, err := grpcservice.userservice.GetProfileByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &grpcrepo.Profile{
		Id:        profile.ID,
		UserID:    profile.UserID,
		Name:      profile.Name,
		Avatar:    profile.Avatar,
		Bio:       profile.Bio,
		CreatedAt: profile.CreatedAt.String(),
	}, nil
}
func (grpcservice *gRPCService) GetUserByEmail(ctx context.Context, req *grpcrepo.GetUserByEmailRequest) (*grpcrepo.User, error) {
	user, err := grpcservice.userservice.GetUserByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	return &grpcrepo.User{
		Id:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		RoleID:    int32(user.RoleID),
		CreatedAt: user.CreatedAt.String(),
	}, nil
}
func (grpcservice *gRPCService) GetUserByID(ctx context.Context, req *grpcrepo.GetUserByIDRequest) (*grpcrepo.User, error) {
	user, err := grpcservice.userservice.GetUserByID(req.Id)

	if err != nil {
		return nil, err
	}

	return &grpcrepo.User{
		Id:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		RoleID:    int32(user.RoleID),
		CreatedAt: user.CreatedAt.String(),
	}, nil
}

func RunGRPCServer(cfg *config.Config, svc *service.UserService) {
	lis, err := net.Listen("tcp", cfg.GRPCPort)
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

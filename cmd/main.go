package main

import (
	"servicerepository/config"
	"servicerepository/internal/grpc"
	"servicerepository/internal/repository"
	"servicerepository/internal/service"

	_ "github.com/lib/pq"
)

func main() {

	cfg := config.LoadConfig()

	db := repository.ConnectToSql(cfg)

	repo := repository.NewPostgresUserRepository(db)
	svc := service.NewUserService(repo)

	grpc.RunGRPCServer(cfg, svc)
}

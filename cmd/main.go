package main

import (
	"servicerepository/internal/grpc"
	"servicerepository/internal/repository"
	"servicerepository/internal/service"

	_ "github.com/lib/pq"
)

func main() {

	db := repository.ConnectToSql()

	repo := repository.NewPostgresUserRepository(db)
	svc := service.NewUserService(repo)

	grpc.RunGRPCServer(svc)

	//http.HandleFunc("/register", handler.Register)

	//log.Println("Server running on :8080")
	//http.ListenAndServe(":8080", nil)
}

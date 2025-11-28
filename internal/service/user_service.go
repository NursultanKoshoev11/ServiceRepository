package service

import (
	"servicerepository/internal/models"
	"servicerepository/internal/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (service *UserService) Register(name, email, password string) (*models.User, error) {
	user := &models.User{Name: name, Email: email, Password: password}
	err := service.repository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	return service.repository.GetByEmail(email)
}

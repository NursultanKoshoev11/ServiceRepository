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

func (service *UserService) DeleteUserByID(user_id int64) error {
	return service.repository.DeleteUserByID(user_id)
}
func (service *UserService) DeleteUserByEmail(email string) error {
	return service.repository.DeleteUserByEmail(email)
}
func (service *UserService) DeleteProfileByUserID(user_id int64) error {
	return service.repository.DeleteProfileByUserID(user_id)
}

func (service *UserService) CreateUser(email, password string, role_id models.RoleType) (*models.User, error) {
	user := &models.User{
		Email:    email,
		Password: password,
		RoleID:   role_id,
	}
	err := service.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	err = service.repository.CreateProfile(user.ID)

	if err != nil {
		err = service.repository.DeleteUserByID(user.ID)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func (service *UserService) CreateRole(roleType models.RoleType, roleName string) error {
	return service.repository.CreateRole(roleType, roleName)
}
func (service *UserService) GetProfileByID(user_id int64) (*models.Profile, error) {
	return service.repository.GetProfileByUserID(user_id)
}
func (service *UserService) GetProfileByEmail(email string) (*models.Profile, error) {
	return service.repository.GetProfileByEmail(email)
}

func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	return service.repository.GeUserByEmail(email)
}

func (service *UserService) GetUserByID(ID int64) (*models.User, error) {
	return service.repository.GetUserByID(ID)
}

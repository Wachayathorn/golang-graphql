package service

import (
	"github.com/wachayathorn/golang-graphql/graph/model"
	"github.com/wachayathorn/golang-graphql/repository"
)

type UserServiceInterface interface {
	CreateUser(data *model.User) (*model.User, error)
	GetUsers() ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
	UpdateUser(data *model.User) (*model.User, error)
	DeleteUser(id int) (*model.Response, error)
	NewUserRepository(repo repository.UserRepositoryInterface) UserServiceInterface
}

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func (s UserService) NewUserRepository(repo repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{userRepo: repo}
}

func (s UserService) CreateUser(data *model.User) (*model.User, error) {
	return s.userRepo.CreateUser(data)
}

func (s UserService) GetUsers() ([]*model.User, error) {
	return s.userRepo.GetUsers()
}

func (s UserService) GetUserById(id int) (*model.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s UserService) UpdateUser(data *model.User) (*model.User, error) {
	return s.userRepo.UpdateUser(data)
}

func (s UserService) DeleteUser(id int) (*model.Response, error) {
	return s.userRepo.DeleteUser(id)
}

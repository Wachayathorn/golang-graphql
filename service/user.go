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
}

type UserServiceForUnittest struct {
	repository.UserRepositoryInterface
}

type UserService struct {
	repository.UserRepository
}

func (s UserService) CreateUser(data *model.User) (*model.User, error) {
	return s.UserRepository.CreateUser(data)
}

func (s UserService) GetUsers() ([]*model.User, error) {
	return s.UserRepository.GetUsers()
}

func (s UserService) GetUserById(id int) (*model.User, error) {
	return s.UserRepository.GetUserById(id)
}

func (s UserService) UpdateUser(data *model.User) (*model.User, error) {
	return s.UserRepository.UpdateUser(data)
}

func (a UserService) DeleteUser(id int) (*model.Response, error) {
	return a.UserRepository.DeleteUser(id)
}

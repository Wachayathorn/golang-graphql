package service

import (
	"github.com/wachayathorn/golang-graphql/graph/model"
	"github.com/wachayathorn/golang-graphql/repository"
)

type UserService struct {
	repository.UserRepositoryInterface
}

func (s UserService) CreateUser(data *model.User) (*model.User, error) {
	return s.UserRepositoryInterface.CreateUser(data)
}

func (s UserService) GetUsers() ([]*model.User, error) {
	return s.UserRepositoryInterface.GetUsers()
}

func (s UserService) GetUserById(id int) (*model.User, error) {
	return s.UserRepositoryInterface.GetUserById(id)
}

func (s UserService) UpdateUser(data *model.User) (*model.User, error) {
	return s.UserRepositoryInterface.UpdateUser(data)
}

func (a UserService) DeleteUser(id int) (*model.Response, error) {
	return a.UserRepositoryInterface.DeleteUser(id)
}

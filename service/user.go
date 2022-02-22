package service

import (
	"github.com/wachayathorn/golang-graphql/graph/model"
	"github.com/wachayathorn/golang-graphql/repository"
)

func CreateUser(data *model.User) (*model.User, error) {
	return repository.CreateUser(data)
}

func GetUsers() ([]*model.User, error) {
	return repository.GetUser()
}

func GetUserById(id int) (*model.User, error) {
	return repository.GetUserById(id)
}

func UpdateUser(data *model.User) (*model.User, error) {
	return repository.UpdateUser(data)
}

func DeleteUser(id int) (*model.Response, error) {
	return repository.DeleteUser(id)
}

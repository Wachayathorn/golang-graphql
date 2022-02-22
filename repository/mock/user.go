package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/wachayathorn/golang-graphql/graph/model"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m UserRepositoryMock) CreateUser(data *model.User) (*model.User, error) {
	args := m.Called(data)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m UserRepositoryMock) GetUsers() ([]*model.User, error) {
	args := m.Called()
	users := []*model.User{}
	users = append(users, &model.User{ID: 1, Firstname: "John", Lastname: "Doe", Username: "johndoe"})
	return users, args.Error(1)
}

func (m UserRepositoryMock) GetUserById(id int) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m UserRepositoryMock) UpdateUser(data *model.User) (*model.User, error) {
	args := m.Called(data)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m UserRepositoryMock) DeleteUser(id int) (*model.Response, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Response), args.Error(1)
}

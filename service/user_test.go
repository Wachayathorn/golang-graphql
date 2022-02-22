package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wachayathorn/golang-graphql/graph/model"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (r UserRepositoryMock) GetUsers() ([]*model.User, error) {
	users := []*model.User{}
	users = append(users, &model.User{ID: 1, Firstname: "John", Lastname: "Doe", Username: "johndoe"})
	return users, nil
}

func TestService_GetUsers(t *testing.T) {
	mock := UserRepositoryMock{}
	mock.On("GetUsers").Return([]*model.User{}, nil)

	ser := UserService{mock}
	users, _ := ser.GetUsers()
	for i := range users {
		assert.Equal(t, users[i].ID, 1, "GET USERS IS BE PASSED.")
	}
}

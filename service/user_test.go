package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wachayathorn/golang-graphql/graph/model"
	"github.com/wachayathorn/golang-graphql/repository/mock"
)

func TestService_GetUsers(t *testing.T) {
	mockFunction := mock.UserRepositoryMock{}
	mockFunction.On("GetUsers").Return([]*model.User{}, nil)

	ser := UserService{mockFunction}
	users, _ := ser.GetUsers()
	for i := range users {
		assert.Equal(t, users[i].ID, 1, "GET USERS IS BE PASSED.")
	}
}

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wachayathorn/golang-graphql/graph/model"
	"github.com/wachayathorn/golang-graphql/repository/mock"
)

func TestService_CreateUser(t *testing.T) {
	mockFunction := mock.UserRepositoryMock{}
	mockFunction.On("CreateUser", &model.User{Firstname: "test", Lastname: "test", Username: "test"}).Return(&model.User{}, nil)

	service := UserServiceForUnittest{mockFunction}
	user, _ := service.CreateUser(&model.User{Firstname: "test", Lastname: "test", Username: "test"})
	assert.Equal(t, user.ID, 0, "should be equal")
}

// func TestService_GetUsers(t *testing.T) {
// 	mockFunction := mock.UserRepositoryMock{}
// 	mockFunction.On("GetUsers").Return([]*model.User{}, nil)

// 	service := UserServiceForUnittest{mockFunction}
// 	users, _ := service.GetUsers()
// 	for i := range users {
// 		assert.Equal(t, users[i].ID, 1, "GET USERS IS BE PASSED.")
// 	}
// }

// func TestService_GetUserById(t *testing.T) {
// 	mockFunction := mock.UserRepositoryMock{}
// 	mockFunction.On("GetUserById").Return(&model.User{}, nil)

// 	service := UserServiceForUnittest{mockFunction}
// 	user, _ := service.GetUserById(1)
// 	assert.Equal(t, user.ID, 1, "GET USER BY ID IS BE PASSED.")
// 	fmt.Print(user)
// }

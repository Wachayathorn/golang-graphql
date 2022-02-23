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

	service := UserService{mockFunction}
	user, _ := service.CreateUser(&model.User{Firstname: "test", Lastname: "test", Username: "test"})
	assert.Equal(t, user.ID, 0, "CREATE USER PASSED.")
}

func TestService_GetUsers(t *testing.T) {
	mockFunction := mock.UserRepositoryMock{}
	mockFunction.On("GetUsers").Return([]*model.User{}, nil)

	service := UserService{mockFunction}
	users, _ := service.GetUsers()
	for i := range users {
		assert.Equal(t, users[i].ID, 1, "GET USERS IS BE PASSED.")
	}
}

func TestService_GetUserById(t *testing.T) {
	mockFunction := mock.UserRepositoryMock{}
	mockFunction.On("GetUserById", 1).Return(&model.User{}, nil)

	service := UserService{mockFunction}
	user, _ := service.GetUserById(1)
	assert.Equal(t, user.ID, 1, "GET USER BY ID IS BE PASSED.")
}

func TestService_UpdateUser(t *testing.T) {
	mockFunction := mock.UserRepositoryMock{}
	mockFunction.On("UpdateUser", &model.User{ID: 1, Firstname: "test", Lastname: "test", Username: "test"}).Return(&model.User{}, nil)

	service := UserService{mockFunction}
	user, _ := service.UpdateUser(&model.User{ID: 1, Firstname: "test", Lastname: "test", Username: "test"})
	assert.Equal(t, user.ID, 1, "UPDATE USER IS BE PASSED.")
}

func TestService_DeleteUser(t *testing.T) {
	mockFunction := mock.UserRepositoryMock{}
	mockFunction.On("DeleteUser", 1).Return(&model.Response{}, nil)

	service := UserService{mockFunction}
	response, _ := service.DeleteUser(1)
	assert.Equal(t, response.Success, true, "DELETE USER IS BE PASSED.")
}

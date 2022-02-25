package graph

import (
	"github.com/wachayathorn/golang-graphql/repository"
	"github.com/wachayathorn/golang-graphql/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service.UserService
}

func (r Resolver) InitUserService() service.UserServiceInterface {
	service := r.UserService.NewUserRepository(repository.UserRepository{})
	return service
}

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/wachayathorn/golang-graphql/graph/generated"
	"github.com/wachayathorn/golang-graphql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, firstname string, lastname string, username string) (*model.User, error) {
	return r.UserServiceInterface.CreateUser(&model.User{Firstname: firstname, Lastname: lastname, Username: username})
}

func (r *mutationResolver) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return r.UserServiceInterface.GetUserById(id)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, firstname string, lastname string, username string) (*model.User, error) {
	return r.UserServiceInterface.UpdateUser(&model.User{ID: id, Firstname: firstname, Lastname: lastname, Username: username})
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.Response, error) {
	return r.UserServiceInterface.DeleteUser(id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UserServiceInterface.GetUsers()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

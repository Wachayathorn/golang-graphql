package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/wachayathorn/golang-graphql/graph/generated"
	"github.com/wachayathorn/golang-graphql/graph/model"
	"github.com/wachayathorn/golang-graphql/service"
)

func (r *mutationResolver) CreateUser(ctx context.Context, firstname string, lastname string, username string) (*model.User, error) {
	return service.CreateUser(ctx, &model.User{Firstname: firstname, Lastname: lastname, Username: username})
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := service.GetUsers()
	return users, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

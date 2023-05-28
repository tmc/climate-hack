package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"go-thorium/graph/model"
)

// InformNonBeliver is the resolver for the informNonBeliver field.
func (r *mutationResolver) InformNonBeliver(ctx context.Context, input model.InformNonBeliverInput) (*model.InformNonBeliverPayload, error) {
	_id++
	fact := r.service.GetThoriumFact(ctx)
	r.service.SendSMS(ctx, input.Phone, fact)
	return &model.InformNonBeliverPayload{
		User: &model.User{
			ID:    fmt.Sprint(_id),
			Phone: input.Phone,
		},
		Message: fact,
	}, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return &model.User{
		ID:    "1",
		Phone: "+123456789",
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var _id = 0

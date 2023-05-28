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
	fact := r.service.GetInitialThoriumFact(ctx)
	// TODO: we should probably have an easy way to disable sending SMS
	if err := r.service.SendSMS(ctx, input.Phone, fact); err != nil {
		return nil, err
	}
	r.service.Repository.CreateUser(input.Phone)
	c, err := r.service.Repository.CreateConversation(input.Phone)
	if err != nil {
		return nil, fmt.Errorf("could not create conversation: %w", err)
	}
	botRole := model.MessageRoleBot
	if err := r.service.Repository.AddToConversation(input.Phone, c.ID, model.Message{
		Body: fact,
		Role: &botRole,
	}); err != nil {
		return nil, fmt.Errorf("could not add message to conversation: %w", err)
	}
	c, _ = r.service.Repository.GetConversation(input.Phone, c.ID)
	return &model.InformNonBeliverPayload{
		Conversation: c,
	}, nil
}

// StartConversation is the resolver for the startConversation field.
func (r *mutationResolver) StartConversation(ctx context.Context, input model.StartConversationInput) (*model.StartConversationPayload, error) {
	return nil, fmt.Errorf("not implemented: StartConversation - startConversation")
}

// ContinueConversation is the resolver for the continueConversation field.
func (r *mutationResolver) ContinueConversation(ctx context.Context, input model.ContinueConversationInput) (*model.ContinueConversationPayload, error) {
	err := r.service.Repository.AddToConversation(input.UserID, input.ConversationID, model.Message{
		Body: input.Body,
	})
	if err != nil {
		return nil, fmt.Errorf("could not add message to conversation: %w", err)
	}
	c, _ := r.service.Repository.GetConversation(input.UserID, input.ConversationID)

	return &model.ContinueConversationPayload{
		Conversation: c,
	}, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return &model.User{
		ID:    "1",
		Phone: "+123456789",
	}, nil
}

// MessageAdded is the resolver for the messageAdded field.
func (r *subscriptionResolver) MessageAdded(ctx context.Context, conversationID string) (<-chan *model.Message, error) {
	return r.service.Repository.SubscribeToConversation(conversationID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

var _id = 0

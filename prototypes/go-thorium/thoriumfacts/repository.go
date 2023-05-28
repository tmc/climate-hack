package thoriumfacts

import (
	"fmt"
	"go-thorium/graph/model"
)

// Repository is the interface for the repository
type Repository interface {
	CreateUser(userID string) error
	GetUser(userID string) (*model.User, error)
	UpdateUser(userID string, user model.User) error
	DeleteUser(userID string) error

	CreateConversation(userID string) (*model.Conversation, error)
	GetConversation(userID string, conversationID string) (*model.Conversation, error)
	GetConversations(userID string) ([]*model.Conversation, error)
	AddToConversation(userID string, conversationID string, message model.Message) error
	DeleteConversation(userID string, conversationID string) error

	SubscribeToConversation(conversationID string) (<-chan *model.Message, error)
}

type inMemoryRepository struct {
	// map of userID to conversationID to conversation
	users         map[string]*model.User
	conversations map[string]*model.Conversation

	// map of conversationID to channel
	conversationSubscriptions map[string]chan *model.Message
}

func newInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		users:                     make(map[string]*model.User),
		conversations:             make(map[string]*model.Conversation),
		conversationSubscriptions: make(map[string]chan *model.Message),
	}
}

func (i *inMemoryRepository) CreateUser(userID string) error {
	i.users[userID] = &model.User{
		ID: userID,
	}
	return nil
}

func (i *inMemoryRepository) GetUser(userID string) (*model.User, error) {
	return i.users[userID], nil
}

func (i *inMemoryRepository) UpdateUser(userID string, user model.User) error {
	i.users[userID] = &user
	return nil
}

func (i *inMemoryRepository) DeleteUser(userID string) error {
	delete(i.users, userID)
	return nil
}

func (i *inMemoryRepository) CreateConversation(userID string) (*model.Conversation, error) {
	u, err := i.GetUser(userID)
	if err != nil {
		return nil, err
	}
	c := &model.Conversation{
		ID: fmt.Sprintf("%v-%d", userID, len(u.Conversations)),
	}
	u.Conversations = append(u.Conversations, c)
	i.conversations[c.ID] = c
	// set up channel:
	i.conversationSubscriptions[c.ID] = make(chan *model.Message)
	return c, nil
}

func (i *inMemoryRepository) GetConversation(userID string, conversationID string) (*model.Conversation, error) {
	c, ok := i.conversations[conversationID]
	if !ok {
		return nil, fmt.Errorf("conversation %v not found", conversationID)
	}
	return c, nil
}

func (i *inMemoryRepository) GetConversations(userID string) ([]*model.Conversation, error) {
	u, err := i.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return u.Conversations, nil
}

func (i *inMemoryRepository) AddToConversation(userID string, conversationID string, message model.Message) error {
	c, err := i.GetConversation(userID, conversationID)
	if err != nil {
		return err
	}
	c.Messages = append(c.Messages, &message)
	// send to channel:
	select {
	case i.conversationSubscriptions[c.ID] <- &message:
	default:
	}
	return nil
}

func (i *inMemoryRepository) DeleteConversation(userID string, conversationID string) error {
	c, err := i.GetConversation(userID, conversationID)
	if err != nil {
		return err
	}
	delete(i.conversations, c.ID)
	// todo: find and remove from user convo slice
	return nil

}

func (i *inMemoryRepository) SubscribeToConversation(conversationID string) (<-chan *model.Message, error) {
	ch, ok := i.conversationSubscriptions[conversationID]
	if !ok {
		return nil, fmt.Errorf("conversation %v not found", conversationID)
	}
	return ch, nil
}

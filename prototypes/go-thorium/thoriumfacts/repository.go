package thoriumfacts

import (
	"context"
	"encoding/json"
	"fmt"
	"go-thorium/graph/model"

	"github.com/redis/go-redis/v9"
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

type redisRepository struct {
	client                    *redis.Client
	context                   context.Context
	conversationSubscriptions map[string]chan *model.Message
}

func newRedisRepository() *redisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &redisRepository{
		client:                    client,
		context:                   context.Background(),
		conversationSubscriptions: make(map[string]chan *model.Message),
	}
}

func (r *redisRepository) CreateUser(userID string) error {
	user := &model.User{
		ID: userID,
	}
	userBytes, _ := json.Marshal(user)
	return r.client.Set(r.context, "user:"+userID, userBytes, 0).Err()
}

func (r *redisRepository) GetUser(userID string) (*model.User, error) {
	val, err := r.client.Get(r.context, "user:"+userID).Result()
	if err != nil {
		return nil, err
	}

	user := &model.User{}
	err = json.Unmarshal([]byte(val), user)

	return user, err
}

func (r *redisRepository) UpdateUser(userID string, user model.User) error {
	userBytes, _ := json.Marshal(user)
	return r.client.Set(r.context, "user:"+userID, userBytes, 0).Err()
}

func (r *redisRepository) DeleteUser(userID string) error {
	return r.client.Del(r.context, "user:"+userID).Err()
}

func (r *redisRepository) CreateConversation(userID string) (*model.Conversation, error) {
	user, err := r.GetUser(userID)
	if err != nil {
		return nil, err
	}

	conversation := &model.Conversation{
		ID: fmt.Sprintf("%v-%d", userID, len(user.Conversations)),
	}
	conversationBytes, err := json.Marshal(conversation)
	if err != nil {
		return nil, err
	}

	err = r.client.Set(r.context, "conversation:"+conversation.ID, conversationBytes, 0).Err()
	if err != nil {
		return nil, err
	}

	// update user's conversations
	user.Conversations = append(user.Conversations, conversation)
	err = r.UpdateUser(userID, *user)
	if err != nil {
		return nil, err
	}

	return conversation, nil
}

func (r *redisRepository) GetConversation(userID string, conversationID string) (*model.Conversation, error) {
	val, err := r.client.Get(r.context, "conversation:"+conversationID).Result()
	if err != nil {
		return nil, err
	}

	conversation := &model.Conversation{}
	err = json.Unmarshal([]byte(val), conversation)
	if err != nil {
		return nil, err
	}

	return conversation, err
}

func (r *redisRepository) GetConversations(userID string) ([]*model.Conversation, error) {
	user, err := r.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user.Conversations, nil
}

func (r *redisRepository) AddToConversation(userID string, conversationID string, message model.Message) error {

	conversation, err := r.GetConversation(userID, conversationID)
	if err != nil {
		return err
	}

	conversation.Messages = append(conversation.Messages, &message)

	c, _ := json.Marshal(conversation)

	err = r.client.Set(r.context, "conversation:"+conversationID, c, 0).Err()
	if err != nil {
		fmt.Println("AddToConversation:", err)
		return err
	}

	message.Conversation = conversation

	select {
	case r.conversationSubscriptions[conversationID] <- &message:
	default:
	}

	// if you'd like to publish this message to a Pub/Sub channel, you could do so here
	// however, keep in mind that this isn't exactly the same as a Go channel
	messageBytes, _ := json.Marshal(message)
	err = r.client.Publish(r.context, "conversation:"+conversationID, messageBytes).Err()

	return err
}

func (r *redisRepository) DeleteConversation(userID string, conversationID string) error {
	return r.client.Del(r.context, "conversation:"+conversationID).Err()
}

func (r *redisRepository) SubscribeToConversation(conversationID string) (<-chan *model.Message, error) {
	ch, ok := r.conversationSubscriptions[conversationID]
	if !ok {
		ch = make(chan *model.Message)
		r.conversationSubscriptions[conversationID] = ch
	}
	return ch, nil
}

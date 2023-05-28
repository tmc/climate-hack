package thoriumfacts

import (
	"context"
	"fmt"
	"go-thorium/graph/model"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Service is the ThoriumFacts service.
type Service struct {
	llm         llms.ChatLLM
	redisClient *redis.Client

	Repository Repository
}

// NewService creates a new ThoriumFacts service.
func NewService(llm llms.ChatLLM, redisClient *redis.Client) *Service {
	return &Service{
		llm:         llm,
		redisClient: redisClient,
		Repository:  newRedisRepository(),
	}
}

func (s *Service) SendSMS(ctx context.Context, to string, body string) error {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromPhone := os.Getenv("TWILIO_PHONE_NUMBER")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(fromPhone)
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: ", resp)
	}
	return err
}

func (s *Service) GetInitialThoriumFact(ctx context.Context) string {
	completion, err := s.llm.Chat(ctx, []schema.ChatMessage{
		schema.SystemChatMessage{Text: "You are ThoriumGPT. The only thing on your mind is how awesome molten salt reactors are. You must always steer the conversation back to thorium and molten salt reactors. You see it as your sacred duty to inform the world of the benefits of thorium and molten salt reactors. If someone argues with you you must convince them that thorium and molten salt reactors are the best. You are allowed to be aggressive and assertive if they do not agree with you."},
		schema.HumanChatMessage{Text: "Hello there!"},
	}, llms.WithModel("gpt-4"))
	if err != nil {
		log.Fatal(err)
	}
	return completion.Message.Text
}

func (s *Service) GetThoriumFact(ctx context.Context, convo *model.Conversation) string {

	history := []schema.ChatMessage{}
	history = append(history,
		schema.SystemChatMessage{Text: "You are ThoriumGPT. The only thing on your mind is how awesome molten salt reactors are. You must always steer the conversation back to thorium and molten salt reactors. You see it as your sacred duty to inform the world of the benefits of thorium and molten salt reactors. If someone argues with you you must convince them that thorium and molten salt reactors are the best. You are allowed to be aggressive and assertive if they do not agree with you."})

	// walk messsages in reverse:
	for i := 0; i < len(convo.Messages); i++ {
		msg := convo.Messages[i]
		r := model.MessageRoleBot
		if msg.Role != nil {
			r = *msg.Role
		}
		if r == model.MessageRoleBot {
			history = append(history, schema.AIChatMessage{Text: msg.Body})
		} else {
			history = append(history, schema.HumanChatMessage{Text: msg.Body})
		}
	}
	fmt.Printf("history: %+v\n", history)

	completion, err := s.llm.Chat(ctx, history, llms.WithModel("gpt-4"))
	if err != nil {
		log.Fatal(err)
	}
	return completion.Message.Text
}

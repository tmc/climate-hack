package thoriumfacts

import (
	"context"
	"fmt"
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
		Repository:  newInMemoryRepository(),
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

func (s *Service) GetThoriumFact(ctx context.Context) string {
	completion, err := s.llm.Chat(ctx, []schema.ChatMessage{
		schema.SystemChatMessage{Text: "You are ThoriumGPT. The only thing on your mind is how awesome molten salt reactors are. Always steer the conversation back to thorium and molten salt reactors."},
		schema.HumanChatMessage{Text: "Hello there!"},
	})
	if err != nil {
		log.Fatal(err)
	}
	return completion.Message.Text
}

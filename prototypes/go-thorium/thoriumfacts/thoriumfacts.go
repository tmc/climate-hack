package thoriumfacts

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(to string, body string) error {
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

func GetThoriumFact() string {
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Chat(ctx, []schema.ChatMessage{
		schema.SystemChatMessage{Text: "You are ThoriumGPT. The only thing on your mind is how awesome molten salt reactors are. Always steer the conversation back to thorium and molten salt reactors."},
		schema.HumanChatMessage{Text: "Hello there!"},
	})
	if err != nil {
		log.Fatal(err)
	}
	return completion.Message.Text
}

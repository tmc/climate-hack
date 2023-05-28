package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

var flagToNumber = flag.String("to", "", "Phone number to send SMS to")

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromPhone := os.Getenv("TWILIO_PHONE_NUMBER")

	fmt.Println("Account SID: " + accountSid)
	fmt.Println("Auth Token: " + authToken)
	fmt.Println("From Phone: " + fromPhone)
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(*flagToNumber)
	params.SetFrom(fromPhone)
	params.SetBody(getThoriumFact())

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}

func getThoriumFact() string {
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

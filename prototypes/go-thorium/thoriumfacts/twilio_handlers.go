package thoriumfacts

import (
	"context"
	"crypto/rand"
	"fmt"
	"go-thorium/graph/model"
	"net/http"

	"github.com/oklog/ulid"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func ptr(s string) *string {
	return &s
}

func newULID() string {
	return ulid.MustNew(ulid.Now(), rand.Reader).String()
}

// HandleIncomingTwilioSMS handles incoming SMS messages from Twilio.
func (s *Service) HandleIncomingTwilioSMS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("twilio handler called")
	payload := twilioApi.ApiV2010Message{}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload.Body = ptr(r.FormValue("Body"))
	payload.From = ptr(r.FormValue("From"))
	payload.To = ptr(r.FormValue("To"))

	userID := *payload.From
	conversationID := fmt.Sprintf("%s-%s", userID, "0")

	role := model.MessageRoleUser

	s.Repository.AddToConversation(userID, conversationID, model.Message{
		ID:   newULID(),
		Body: *payload.Body,
		Role: &role,
	})

	go func() {
		c, err := s.Repository.GetConversation(userID, conversationID)
		if err != nil {
			fmt.Println(err)
		}
		fact := s.GetThoriumFact(context.TODO(), c)
		botRole := model.MessageRoleBot
		err = s.Repository.AddToConversation(userID, conversationID, model.Message{
			ID:   newULID(),
			Body: fact,
			Role: &botRole,
		})
		if err != nil {
			fmt.Println(err)
		}
	}()

	// unmarshal from request body:
	// s.Repository.GetUser(
	//w.Write([]byte("Hello, World!!"))
	w.WriteHeader(http.StatusOK)
}

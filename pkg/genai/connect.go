package genai

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/mirjalilova/api-gateway_blacklist/internal/config"
	"google.golang.org/api/option"
)

func ConnectGenai(cnf *config.Config) (*genai.ChatSession, error) {
	apiKey := cnf.API_KEY
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY environment variable is not set")
	}

	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	advice := model.StartChat()

	return advice, nil
}
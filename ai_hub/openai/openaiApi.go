package openai

import (
	"agi-backend/models"
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

var (
	client *openai.Client
	// model_name = []string{"gpt-4o", "gpt-3.5-turbo", "gpt-4o-mini"}
)

type OpenaiApi struct{}

func convertConversationsToChatCompletionMessages(conversations []models.Conversation) []openai.ChatCompletionMessage {
	var chatMessages []openai.ChatCompletionMessage

	for _, convo := range conversations {
		chatMessage := openai.ChatCompletionMessage{
			Role:    convo.Role,
			Content: convo.Content,
		}
		chatMessages = append(chatMessages, chatMessage)
	}

	return chatMessages
}

func (*OpenaiApi) Chat(conversation []models.Conversation, model string, maxToken uint, Temperature float32) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: convertConversationsToChatCompletionMessages(conversation),
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (*OpenaiApi) Model() {}

func Connect() {
	config := openai.DefaultConfig(os.Getenv("OPENAI_API_KEY"))
	config.BaseURL = os.Getenv("OPENAI_BASE_URL")
	client = openai.NewClientWithConfig(config)
}

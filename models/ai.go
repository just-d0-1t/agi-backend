package models

type Ai interface {
	// Chat use to chat with ai
	Chat(conversation []Conversation, model string, maxToken uint, Temperature float32) (string, error)
	// Model returns selectable model type
	Model()
}

type Conversation struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

package ai_hub

import (
	"agi-backend/ai_hub/openai"
	"agi-backend/models"
)

func FetchAi(aiType string) models.Ai {
	if aiType == "openai" {
		return &openai.OpenaiApi{}
	}

	// default
	return &openai.OpenaiApi{}
}

func InitAI() {
	openai.Connect()
}

func GetAbstract(message string) string {
	return "hello world"
}

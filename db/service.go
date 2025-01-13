package db

import (
	"agi-backend/models"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

func FindUserByName(name string) (*User, error) {
	var user User
	tx := DB.Where("username = ?", name).First(&user)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &user, nil
}

func FindUserByID(id uint) (*User, error) {
	var user User
	tx := DB.First(&user, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &user, nil
}

func SaveUser(user *User) (uint, error) {
	tx := DB.Save(user)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return user.ID, nil
}

func FindAgentByName(name string) int {
	return 1
}

func FindAgentByID(id uint) (*Agent, error) {
	var agent Agent
	tx := DB.First(&agent, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &agent, nil
}

func SaveAgent(agent *Agent) (uint, error) {
	tx := DB.Save(agent)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return agent.ID, nil
}

func NewFaq(agent *Agent, abstract, message string) (uint, []models.Conversation, error) {
	// 拼接对话
	conversation := make([]models.Conversation, 0)
	conversation = append(conversation, models.Conversation{
		Role:    "system",
		Content: agent.Prompt,
	})
	conversation = append(conversation, models.Conversation{
		Role:    "user",
		Content: message,
	})

	conversationMarshaled, _ := json.Marshal(conversation)

	faq := Faq{
		Abstract:     abstract,
		Conversation: string(conversationMarshaled),
	}

	// 存到faq表中，获取到faqID
	tx := DB.Save(&faq)

	// 4. 检查错误并获取插入的 ID
	if tx.Error != nil {
		return 0, nil, fmt.Errorf("Failed to insert data: %v", tx.Error)
	}

	// 保存对话id
	var faqs []uint
	if err := json.Unmarshal([]byte(agent.Faqs), &faqs); err != nil {
		return 0, nil, fmt.Errorf("Failed to insert data: %v", err)
	}
	faqs = append(faqs, faq.ID)
	faqsMarshaled, _ := json.Marshal(faqs)
	agent.Faqs = string(faqsMarshaled)

	return faq.ID, conversation, nil
}

func GetFaq(faqID uint) ([]models.Conversation, error) {
	var faq Faq
	tx := DB.First(&faq, faqID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var conversation []models.Conversation
	if err := json.Unmarshal([]byte(faq.Conversation), &conversation); err != nil {
		return nil, err
	}

	return conversation, nil
}

func SaveFaq(faqID uint, conversation []models.Conversation) error {
	conversationMarshaled, _ := json.Marshal(conversation)
	faq := Faq{
		ID:           faqID,
		Conversation: string(conversationMarshaled),
	}
	tx := DB.Save(&faq)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

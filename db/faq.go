package db

import (
	"agi-backend/models"
	"encoding/json"
)

func NewFaq(agent *Agent, abstract string, conversation []models.Conversation) error {
	conversationMarshaled, _ := json.Marshal(conversation)

	faq := Faq{
		Abstract:     abstract,
		Conversation: string(conversationMarshaled),
	}

	// 存到faq表中，获取到faqID
	tx := DB.Save(&faq)

	// 4. 检查错误并获取插入的 ID
	if tx.Error != nil {
		return tx.Error
	}

	// 保存对话id
	var faqs []uint
	if err := json.Unmarshal([]byte(agent.Faqs), &faqs); err != nil {
		return err
	}
	faqs = append(faqs, faq.ID)
	faqsMarshaled, _ := json.Marshal(faqs)
	agent.Faqs = string(faqsMarshaled)

	tx = DB.Save(agent)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
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

package db

import (
	"agi-backend/models"
	"encoding/json"
)

func NewFaq(agent *Agent, abstract string, conversation []models.Conversation) (uint, error) {
	conversationMarshaled, _ := json.Marshal(conversation)

	faq := Faq{
		Conversation: string(conversationMarshaled),
	}

	// 存到faq表中，获取到faqID
	tx := DB.Save(&faq)

	// 4. 检查错误并获取插入的 ID
	if tx.Error != nil {
		return 0, tx.Error
	}

	// 保存对话id
	agentFaq := AgentFaq{
		AgentID:  agent.ID,
		FaqID:    faq.ID,
		Abstract: abstract,
	}

	tx = DB.Save(&agentFaq)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return faq.ID, nil
}

func GetFaqByAgentID(agentID uint) (*[]AgentFaq, error) {
	// 查询所有 AgentID = 1 的记录
	var faqRecords []AgentFaq
	tx := DB.Where("agent_id = ?", agentID).Find(&faqRecords)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &faqRecords, nil
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

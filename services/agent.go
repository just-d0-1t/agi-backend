package services

import (
	"agi-backend/ai_hub"
	"agi-backend/db"
	"agi-backend/models"
	"agi-backend/utils"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func CreateAgent(c *gin.Context) {
	var agent db.Agent
	if err := c.ShouldBindJSON(&agent); err != nil {
		utils.ResponseError(c, err.Error())
		return
	}
	if agent.Name == "" {
		utils.ResponseError(c, "Agent name is required.")
		return
	}
	if agent.Faqs == "" {
		emptyJSON := []uint{}
		marshaledJSON, _ := json.Marshal(emptyJSON)
		agent.Faqs = string(marshaledJSON)
	}

	// 如果agent存在，则更新，否则创建agent
	agentID, err := db.SaveAgent(&agent)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	utils.ResponseSuccess(c, agentID)
}

func GetAgent(c *gin.Context) {
	var AgentRequest db.Agent
	if err := c.ShouldBindJSON(&AgentRequest); err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	agent, err := db.FindAgentByID(AgentRequest.ID)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	agentMarshaled, _ := json.Marshal(agent)
	utils.ResponseSuccess(c, agentMarshaled)
}

type FaqRequest struct {
	FaqID   uint   `json:"faq_id"`
	AgentID uint   `json:"agent_id"`
	Message string `json:"message"`
}

func Faq(c *gin.Context) {
	var faqRequest FaqRequest
	// 获取到agentID、message、faqID
	if err := c.ShouldBindJSON(&faqRequest); err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	// 获取agent信息
	agent, err := db.FindAgentByID(faqRequest.AgentID)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	// 获取ai
	ai := ai_hub.FetchAi(agent.AiType)

	// 保存聊天记录，保存faqID至agent
	var conversation []models.Conversation

	// 如果faqID是0，则创建一个新的对话ID
	// 让ai根据首句生成一个摘要作为对话的标题
	if faqRequest.FaqID == 0 {
		abstract := ai_hub.GetAbstract(faqRequest.Message)
		faqRequest.FaqID, conversation, err = db.NewFaq(agent, abstract, faqRequest.Message)
		if err != nil {
			utils.ResponseError(c, err.Error())
			return
		}
	} else { // 根据对话ID获取聊天记录 TODO
		conversation, err = db.GetFaq(faqRequest.FaqID)
		if err != nil {
			utils.ResponseError(c, err.Error())
			return
		}
		conversation = append(conversation, models.Conversation{
			Role:    "user",
			Content: faqRequest.Message,
		})
	}

	// 获取ai的回复 TODO
	response, err := ai.Chat(conversation, agent.ModelName, agent.MaxToken, agent.Temperature)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	conversation = append(conversation, models.Conversation{
		Role:    "assistant",
		Content: response,
	})

	// 保存到数据库
	if err := db.SaveFaq(faqRequest.FaqID, conversation); err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	// 然后响应回复给客户端
	utils.ResponseSuccess(c, response)
}

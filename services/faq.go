package services

import (
	"agi-backend/db"
	"agi-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetFaqs(c *gin.Context) {
	var agentRequest AgentRequest
	if err := c.ShouldBindJSON(&agentRequest); err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	faqs, err := db.GetFaqByAgentID(agentRequest.ID)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	utils.ResponseSuccess(c, faqs)
}

type GetFaqRequest struct {
	ID uint `json:"id"`
}

func GetFaq(c *gin.Context) {
	var faqRequest GetFaqRequest
	if err := c.ShouldBindJSON(&faqRequest); err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	faqs, err := db.GetFaq(faqRequest.ID)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	// faqsMarshaled, _ := json.Marshal(faqs)
	utils.ResponseSuccess(c, faqs)
}

package db

import (
	"encoding/json"
)

func FindAgentByName(name string) int {
	return 1
}

func FindAgentByID(id uint) (*Agent, error) {
	var agent Agent
	tx := DB.First(&agent, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &agent, nil
}

func SaveAgent(userID uint, agent *Agent) (uint, error) {
	user, err := FindUserByID(userID)
	if err != nil {
		return 0, err
	}

	var agents []uint
	if err := json.Unmarshal([]byte(user.AgentID), &agents); err != nil {
		return 0, err
	}

	tx := DB.Save(agent)
	if tx.Error != nil {
		return 0, tx.Error
	}

	agents = append(agents, agent.ID)
	agentMashaled, _ := json.Marshal(agents)
	user.AgentID = string(agentMashaled)

	tx = DB.Save(user)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return agent.ID, nil
}

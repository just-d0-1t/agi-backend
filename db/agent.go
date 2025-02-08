package db

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

func FindAgentByUserID(id uint) []UserAgent {
	var agents []UserAgent
	// 查询 UserAgent 表中与 UserID 相关的记录
	DB.Where("user_id = ?", id).Find(&agents)
	return agents
}

func SaveAgent(userID uint, agent *Agent) (uint, error) {
	user, err := FindUserByID(userID)
	if err != nil {
		return 0, err
	}

	tx := DB.Save(agent)
	if tx.Error != nil {
		return 0, tx.Error
	}

	user_agent := UserAgent{
		UserID:    user.ID,
		AgentID:   agent.ID,
		AgentName: agent.Name,
	}

	tx = DB.Save(user_agent)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return agent.ID, nil
}

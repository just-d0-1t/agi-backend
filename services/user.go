package services

import (
	"agi-backend/db"
	"agi-backend/services/middleware"
	"agi-backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var userInfo db.User

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	if len(userInfo.Username) == 0 || len(userInfo.Password) == 0 {
		utils.ResponseError(c, fmt.Errorf("userinfo is illegel").Error())
		return
	}

	// 将用户注册到数据库表中
	userExist, err := db.FindUserByName(userInfo.Username)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}
	if userExist != nil {
		utils.ResponseError(c, fmt.Errorf("username is conflict").Error())
		return
	}

	userID, err := db.SaveUser(&userInfo)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	agent := db.Agent{
		Name:        "聊天机器人",
		MaxToken:    1000,
		Temperature: 0.1,
	}

	createAgent(userID, &agent)

	utils.ResponseSuccess(c, userID)
}

type AgentInfo struct {
	AgentID   uint   `json:"agent_id"`
	AgentName string `json:"agent_name"`
}

type loginResponse struct {
	Token    string      `json:"token"`
	ID       uint        `json:"ID"`
	Username string      `json:"username"`
	Agents   []AgentInfo `json:"agents"`
}

// // 用户登陆
func Login(c *gin.Context) {
	var userInfo db.User

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	// 查询用户
	user, err := db.FindUserByName(userInfo.Username)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	if user == nil {
		utils.ResponseError(c, "user is unexist")
		return
	}

	if userInfo.Password != user.Password {
		utils.ResponseError(c, fmt.Errorf("password is uncorrect").Error())
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	loginRes := loginResponse{
		Token:    token,
		ID:       user.ID,
		Username: user.Username,
	}

	agents := db.FindAgentByUserID(user.ID)

	// 查询机器人信息
	for _, agent := range agents {
		loginRes.Agents = append(loginRes.Agents, AgentInfo{
			AgentID:   agent.AgentID,
			AgentName: agent.AgentName,
		})
	}

	utils.ResponseSuccess(c, loginRes)
}

// // 用户权限认证
// func Authorizion(userConfig *configs.UserConf) error {

// }

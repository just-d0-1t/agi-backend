package services

import (
	"agi-backend/db"
	"agi-backend/services/middleware"
	"agi-backend/utils"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	jwtSecret = "agi-backend-jwt-secret"
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
		utils.ResponseError(c, fmt.Errorf("username is conflict.").Error())
		return
	}

	if userInfo.AgentID == "" {
		emptyJSON := []uint{}
		marshaledJSON, _ := json.Marshal(emptyJSON)
		userInfo.AgentID = string(marshaledJSON)
	}

	userID, err := db.SaveUser(&userInfo)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	utils.ResponseSuccess(c, userID)
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

	if userInfo.Password != user.Password {
		utils.ResponseError(c, fmt.Errorf("password is uncorrect.").Error())
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	utils.ResponseSuccess(c, token)
}

// // 用户权限认证
// func Authorizion(userConfig *configs.UserConf) error {

// }

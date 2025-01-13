package services

import (
	"agi-backend/ai_hub"
	"agi-backend/configs"
	"agi-backend/db"
	"agi-backend/services/middleware"

	"github.com/gin-gonic/gin"
)

func InitModel() {
	db.InitDB(configs.GlobalConf.DB)
	ai_hub.InitAI()
}

func SetupRouter() *gin.Engine {
	g := gin.Default()

	g.POST("/register", Register)
	g.POST("/login", Login)
	apiV1 := g.Group("v1")
	apiV1.Use(middleware.AuthMiddleware())
	apiV1.POST("/faq", Faq)
	apiV1.POST("/agent", CreateAgent)
	apiV1.GET("/fetch/agent", GetAgent)

	return g
}

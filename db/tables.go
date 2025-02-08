package db

// 声明各种数据库表

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type DeletedAt sql.NullTime

type Model struct {
	ID        uint `gorm:"primarykey; autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Model
	Username string `gorm:"not null; unique; index; comment:用户名" json:"username"`
	Password string `gorm:"comment:密码" json:"password"`
}

type UserAgent struct {
	UserID    uint   `gorm:"primaryKey" json:"user_id"`
	AgentID   uint   `gorm:"primaryKey" json:"agent_id"`
	AgentName string `gorm:"not null;" json:"agent_name"`
}

type Agent struct {
	Model
	Name        string  `gorm:"not null; index" json:"name"`
	Desc        string  `gorm:"TYPE:varchar(1024);charset:utf8mb4;collate:utf8mb4_unicode_ci" json:"description"`
	Prompt      string  `gorm:"TYPE:varchar(4096);charset:utf8mb4;collate:utf8mb4_unicode_ci" json:"prompt"`
	MaxToken    uint    `gorm:"not null; default: 1024" json:"max_token"`
	Temperature float32 `gorm:"not null; default:0.1" json:"temperature"`
	AiType      string  `gorm:"default:openai" json:"api_type"`
	ModelName   string  `gorm:"default:gpt-4o-mini" json:"model_name"`
}

type AgentKnowledge struct {
	AgentID     uint `gorm:"primarykey" json:"agent_id"`
	KnowledgeID uint `gorm:"primarykey" json:"knowledge_id"`
}

type AgentFaq struct {
	AgentID  uint   `gorm:"primarykey" json:"agent_id"`
	FaqID    uint   `gorm:"primarykey" json:"faq_id"`
	Abstract string `gorm:"TYPE:varchar(100)" json:"abstract"`
}

type Faq struct {
	ID           uint   `gorm:"primarykey; autoIncrement"`
	Conversation string `gorm:"TYPE:json"`
}

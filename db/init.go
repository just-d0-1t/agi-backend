package db

import (
	"agi-backend/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB(dbconf *configs.DBConf) error {
	// 初始化数据库配置
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.Username, dbconf.Password, dbconf.Host, dbconf.Port, dbconf.DBName)

	// 登陆数据库
	// 返回数据库对象
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlDsn,
		DefaultStringSize:         191,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		// Logger:  logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	if err = DB.AutoMigrate(&User{}, &UserAgent{}, &Agent{}, &AgentFaq{}, &Faq{}); err != nil {
		return err
	}

	return err
}

// 对数据库表中的各种操作 ，是否应该封装在db这个模块中？
// 还是说db只提供基础的读写操作，业务逻辑应该在上层？
// 我认为应该是后者，而且数据库对象应该在模块之中

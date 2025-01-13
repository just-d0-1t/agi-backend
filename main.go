package main

import (
	"agi-backend/configs"
	"agi-backend/services"
	"flag"
)

func init() {
	// 配置解析
	var configPath string

	flag.StringVar(&configPath, "c", "./config.toml", "specify a toml file to set config")
	flag.Parse()

	if err := configs.Parse(configPath); err != nil {
		panic(err)
	}
}

func main() {
	// 初始化各模块
	services.InitModel()

	// 注册handler
	app := services.SetupRouter()

	// 监听服务
	app.Run(configs.GlobalConf.Http.ListenAddr)
}

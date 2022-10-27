package main

import (
	"bind/config"
	"bind/public/common"
)

func main() {
	// 加载配置文件到全局配置结构体
	config.InitConfig()

	// 初始化日志
	common.InitLogger()

	// 初始化数据库(mysql)
	common.InitMysql()
}

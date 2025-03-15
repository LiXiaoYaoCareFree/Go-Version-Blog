package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
	"Blog-Server/router"
	"Blog-Server/service/cron_service"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	global.Redis = core.InitRedis()
	global.ESClient = core.EsConnect()
	core.InitIPDB()
	flags.Run()
	core.InitMysqlES()
	// 定时任务
	cron_service.Cron()
	//启动web程序
	router.Run()
}

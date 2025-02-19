// testdata/5.运行日志.go
package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
	"Blog-Server/service/log_service"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	log := log_service.NewRuntimeLog("同步文章数据", log_service.RuntimeDateHour)
	log.SetItem("文章1", 11)
	log.Save()
	log.SetItem("文章2", 12)
	log.Save()
}

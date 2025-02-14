package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	flags.Run()

	//启动web程序

}

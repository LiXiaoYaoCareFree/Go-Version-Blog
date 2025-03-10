package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
	"Blog-Server/service/ai_service"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()

	msg, err := ai_service.Chat("你好")
	fmt.Println(msg, err)

	msgChan, err := ai_service.ChatStream("你好")
	if err != nil {
		fmt.Println(err)
		return
	}
	for s := range msgChan {
		fmt.Println(s)
	}
}

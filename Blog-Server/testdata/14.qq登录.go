// testdata/14.qq登录.go
package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
	"Blog-Server/service/qq_service"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	fmt.Println(qq_service.GetUserInfo("C7C4360BE86B5B91F5CFE3A0ECC6574C"))
}

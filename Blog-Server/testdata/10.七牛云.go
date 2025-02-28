// testdata/10.七牛云.go
package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
	"Blog-Server/service/qiniu_service"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	//url, err := SendFile("uploads/images/头像_0003_26.jpg")
	//fmt.Println(url, err)

	//file, _ := os.Open("uploads/images/头像_0003_26.jpg")
	//
	//url, err := SendReader(file)
	//fmt.Println(url, err)
	fmt.Println(qiniu_service.GenToken())
}

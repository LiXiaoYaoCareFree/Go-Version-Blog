// testdata/6.jwt.go
package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
	"Blog-Server/utils/jwts"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	token, err := jwts.GetToken(jwts.Claims{
		UserID: 1,
		Role:   2,
	})
	fmt.Println(token, err)
	//cls, err := jwts.ParseToken("xx")
	//fmt.Println(cls, err)
}

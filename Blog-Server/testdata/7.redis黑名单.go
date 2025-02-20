// testdata/7.redis黑名单.go
package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
	"Blog-Server/service/redis_service/redis_jwt"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()

	//token, err := jwts.GetToken(jwts.Claims{
	//	UserID: 2,
	//	Role:   1,
	//})
	//fmt.Println(token, err)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsInVzZXJuYW1lIjoiIiwicm9sZSI6MiwiZXhwIjoxNzQwMDMwMDg2LCJpc3MiOiJsaW5neXVhbiJ9.bv6Wl9mCn6M7p_bD1z4sk8SClqeUIhdIN3idVoLjkcc"
	redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)
	blk, ok := redis_jwt.HasTokenBlack(token)
	fmt.Println(blk, ok)
}

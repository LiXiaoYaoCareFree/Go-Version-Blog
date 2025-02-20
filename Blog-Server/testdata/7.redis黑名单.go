// testdata/7.redis黑名单.go
package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
	"Blog-Server/service/redis_service/redis_jwt"
	"Blog-Server/utils/jwts"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()

	token, err := jwts.GetToken(jwts.Claims{
		UserID: 2,
		Role:   1,
	})
	fmt.Println(token, err)
	redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)
	blk, ok := redis_jwt.HasTokenBlack(token)
	fmt.Println(blk, ok)
}

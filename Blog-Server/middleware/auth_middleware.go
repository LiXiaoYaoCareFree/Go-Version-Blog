// middleware/auth_middleware.go
package middleware

import (
	"Blog-Server/common/res"
	"Blog-Server/models/enum"
	"Blog-Server/service/redis_service/redis_jwt"
	"Blog-Server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	blcType, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blcType.Msg(), c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
	return
}

func AdminMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	if claims.Role != enum.AdminRole {
		res.FailWithMsg("权限错误", c)
		c.Abort()
		return
	}
	blcType, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blcType.Msg(), c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
}

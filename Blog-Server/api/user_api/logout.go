package user_api

import (
	"Blog-Server/common/res"
	"Blog-Server/service/redis_service/redis_jwt"
	"github.com/gin-gonic/gin"
)

func (UserApi) LogoutView(c *gin.Context) {
	token := c.Request.Header.Get("token")
	redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)

	res.OkWithMsg("注销成功", c)
}

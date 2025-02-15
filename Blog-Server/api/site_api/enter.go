package site_api

import (
	"Blog-Server/models/enum"
	"Blog-Server/service/log_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	fmt.Println("1")
	log_service.NewLoginSuccess(c, enum.UserPwdLoginType)
	log_service.NewLoginFail(c, enum.UserPwdLoginType, "用户不存在", "lingyuan", "1234")
	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
	return
}

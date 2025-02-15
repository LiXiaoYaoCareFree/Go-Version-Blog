// service/log_service/login_log.go
package log_service

import (
	"Blog-Server/core"
	"Blog-Server/global"
	"Blog-Server/models"
	"Blog-Server/models/enum"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	token := c.GetHeader("token")
	fmt.Println(token)
	// TODO: 通过jwt获取用户id
	userID := uint(1)
	userName := ""
	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		Content:     "",
		UserID:      userID,
		IP:          ip,
		Addr:        addr,
		LoginStatus: true,
		Username:    userName,
		Pwd:         "",
		LoginType:   loginType,
	})
}

func NewLoginFail(c *gin.Context, loginType enum.LoginType, msg string, username string, pwd string) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录失败",
		Content:     msg,
		IP:          ip,
		Addr:        addr,
		LoginStatus: false,
		Username:    username,
		Pwd:         pwd,
		LoginType:   loginType,
	})
}

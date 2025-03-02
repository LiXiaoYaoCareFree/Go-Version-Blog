// api/user_api/qq_login.go
package user_api

import (
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/models"
	"Blog-Server/models/enum"
	"Blog-Server/service/qq_service"
	"Blog-Server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type QQLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

func (UserApi) QQLoginView(c *gin.Context) {
	var cr QQLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	if !global.Config.Site.Login.QQLogin {
		res.FailWithMsg("站点未启用QQ登陆", c)
		return
	}

	info, err := qq_service.GetUserInfo(cr.Code)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "open_id = ?", info.OpenID).Error
	if err != nil {
		// 创建用户
		uname := base64Captcha.RandText(5, "0123456789")
		user = models.UserModel{
			Username:       fmt.Sprintf("b_%s", uname),
			Nickname:       info.Nickname,
			Avatar:         info.Avatar,
			RegisterSource: enum.RegisterQQSourceType,
			OpenID:         info.OpenID,
			Role:           enum.UserRole,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMsg("qq登录失败", c)
			return
		}
	}

	// 颁发token
	token, _ := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})

	res.OkWithData(token, c)
}

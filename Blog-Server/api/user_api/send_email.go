// api/user_api/send_email.go
package user_api

import (
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/models"
	"Blog-Server/models/enum"
	"Blog-Server/service/email_service"
	"Blog-Server/utils/email_store"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type SendEmailRequest struct {
	Type  int8   `json:"type" binding:"oneof=1 2 3"` // 1 注册 2 重置密码 3 绑定邮箱
	Email string `json:"email" binding:"required"`
}

type SendEmailResponse struct {
	EmailID string `json:"emailID"`
}

func (UserApi) SendEmailView(c *gin.Context) {
	var cr SendEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg("站点未启用邮箱注册", c)
		return
	}

	code := base64Captcha.RandText(4, "0123456789")
	id := base64Captcha.RandomId()

	switch cr.Type {
	case 1:
		// 查邮箱是否不存在
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err == nil {
			res.FailWithMsg("该邮箱已使用", c)
			return
		}
		err = email_service.SendRegisterCode(cr.Email, code)
	case 2:
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err != nil {
			res.FailWithMsg("该邮箱不存在", c)
			return
		}
		// 还必须得是邮箱注册的
		if user.RegisterSource != enum.RegisterEmailSourceType {
			res.FailWithMsg("非邮箱注册用户，不能重置密码", c)
			return
		}
		err = email_service.SendResetPwdCode(cr.Email, code)
	case 3:
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err == nil {
			res.FailWithMsg("该邮箱已使用", c)
			return
		}
		err = email_service.SendBIndEmailCode(cr.Email, code)
	}
	if err != nil {
		logrus.Errorf("邮件发送失败 %s", err)
		res.FailWithMsg("邮件发送失败", c)
		return
	}

	email_store.Set(id, cr.Email, code)

	res.OkWithData(SendEmailResponse{
		EmailID: id,
	}, c)
}

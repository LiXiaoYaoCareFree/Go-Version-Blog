// api/user_api/pwd_login.go
package user_api

import (
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"Blog-Server/service/user_service"
	"Blog-Server/utils/jwts"
	"Blog-Server/utils/pwd"
	"github.com/gin-gonic/gin"
)

type PwdLoginRequest struct {
	Val      string `json:"val" binding:"required"` // 有可能是用户名，也可能是邮箱
	Password string `json:"password" binding:"required"`
}

func (UserApi) PwdLoginApi(c *gin.Context) {
	cr := middleware.GetBind[PwdLoginRequest](c)

	if !global.Config.Site.Login.UsernamePwdLogin {
		res.FailWithMsg("站点未启用密码登陆", c)
		return
	}

	var user models.UserModel
	err := global.DB.Take(&user, "(username = ? or email = ?) and password <> ''",
		cr.Val, cr.Val).Error
	if err != nil {
		res.FailWithMsg("用户名密码错误", c)
		return
	}
	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		res.FailWithMsg("用户名密码错误", c)
		return
	}

	// 颁发token
	token, _ := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	user_service.NewUserService(user).UserLogin(c)
	res.OkWithData(token, c)
}

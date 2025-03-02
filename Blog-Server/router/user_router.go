// router/site_router.go
package router

import (
	"Blog-Server/api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	r.POST("user/send_email", middleware.CaptchaMiddleware, app.SendEmailView)
	r.POST("user/email", middleware.EmailVerifyMiddleware, app.RegisterEmailView)
	r.POST("user/qq", app.QQLoginView)
	r.POST("user/login", middleware.CaptchaMiddleware, app.PwdLoginApi)
	r.GET("user/detail", middleware.AuthMiddleware, app.UserDetailView)
	r.GET("user/login", middleware.AuthMiddleware, app.UserLoginListView)
	r.GET("user/base", app.UserBaseInfoView)
	r.PUT("user/password", middleware.AuthMiddleware, app.UpdatePasswordView)
	r.PUT("user/password/reset", middleware.EmailVerifyMiddleware, app.ResetPasswordView)
	r.PUT("user/email/bind", middleware.EmailVerifyMiddleware, middleware.AuthMiddleware, app.BindEmailView)
}

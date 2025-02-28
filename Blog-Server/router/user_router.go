// router/site_router.go
package router

import (
	"Blog-Server/api"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	r.POST("user/send_email", app.SendEmailView)
}

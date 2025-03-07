package router

import (
	"Blog-Server/api"
	"Blog-Server/api/focus_api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func FocusRouter(r *gin.RouterGroup) {
	app := api.App.FocusApi
	r.POST("focus", middleware.AuthMiddleware, middleware.BindJsonMiddleware[focus_api.FocusUserRequest], app.FocusUserView)
	r.GET("focus/my_focus", middleware.AuthMiddleware, middleware.BindQueryMiddleware[focus_api.FocusUserListRequest], app.FocusUserListView)
}

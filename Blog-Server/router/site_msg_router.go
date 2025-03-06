// router/site_router.go
package router

import (
	"Blog-Server/api"
	"Blog-Server/api/site_msg_api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func SiteMsgRouter(r *gin.RouterGroup) {
	app := api.App.SiteMsgApi
	r.GET("site_msg/conf", middleware.AuthMiddleware, app.UserSiteMessageConfView)
	r.PUT("site_msg/conf", middleware.AuthMiddleware, middleware.BindJsonMiddleware[site_msg_api.UserMessageConfUpdateRequest], app.UserSiteMessageConfUpdateView)
	r.GET("site_msg", middleware.AuthMiddleware, middleware.BindQueryMiddleware[site_msg_api.SiteMsgListRequest], app.SiteMsgListView)
	r.POST("site_msg", middleware.AuthMiddleware, middleware.BindJsonMiddleware[site_msg_api.SiteMsgReadRequest], app.SiteMsgReadView)
	r.DELETE("site_msg", middleware.AuthMiddleware, middleware.BindJsonMiddleware[site_msg_api.SiteMsgRemoveRequest], app.SiteMsgRemoveView)
}

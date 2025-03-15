// router/site_router.go
package router

import (
	"Blog-Server/api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	app := api.App.BannerApi
	r.GET("banner", middleware.CacheMiddleware(middleware.NewBannerCacheOption()), app.BannerListView)
	r.POST("banner", middleware.AdminMiddleware, app.BannerCreateView)
	r.PUT("banner/:id", middleware.AdminMiddleware, app.BannerUpdateView)
	r.DELETE("banner", middleware.AdminMiddleware, app.BannerRemoveView)
}

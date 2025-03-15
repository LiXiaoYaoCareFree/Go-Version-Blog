// router/site_router.go
package router

import (
	"Blog-Server/api"
	"Blog-Server/api/data_api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func DataRouter(r *gin.RouterGroup) {
	app := api.App.DataApi
	r.GET("data/sum", middleware.AdminMiddleware, app.SumView)
	r.GET("data/article/year", middleware.AdminMiddleware, app.ArticleYearDataView)
	r.GET("data/computer", middleware.AdminMiddleware, middleware.CacheMiddleware(middleware.NewDataCacheOption()), app.ComputerDataView)
	r.GET("data/growth", middleware.AdminMiddleware, middleware.BindQueryMiddleware[data_api.GrowthDataRequest], app.GrowthDataView)
}

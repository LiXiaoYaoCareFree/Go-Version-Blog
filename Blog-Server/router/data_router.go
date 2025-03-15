// router/site_router.go
package router

import (
	"Blog-Server/api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func DataRouter(r *gin.RouterGroup) {
	app := api.App.DataApi
	r.GET("data/sum", middleware.AdminMiddleware, app.SumView)
}

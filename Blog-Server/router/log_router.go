// router/log_router.go
package router

import (
	"Blog-Server/api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func LogRouter(rr *gin.RouterGroup) {
	app := api.App.LogApi
	r := rr.Group("").Use(middleware.AdminMiddleware)
	r.GET("logs", app.LogListView)
	r.GET("logs/:id", app.LogReadView)
	r.DELETE("logs", app.LogRemoveView)
}

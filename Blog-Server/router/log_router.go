// router/log_router.go
package router

import (
	"Blog-Server/api"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi
	r.GET("logs", app.LogListView)
}

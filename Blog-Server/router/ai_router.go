package router

import (
	"Blog-Server/api"
	"Blog-Server/api/ai_api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func AiRouter(r *gin.RouterGroup) {
	app := api.App.AiApi
	r.POST("ai/analysis", middleware.AuthMiddleware, middleware.BindJsonMiddleware[ai_api.ArticleAnalysisRequest], app.ArticleAnalysisView)
}

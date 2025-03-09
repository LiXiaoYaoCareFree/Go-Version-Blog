package router

import (
	"Blog-Server/api"
	"Blog-Server/api/search_api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func SearchRouter(r *gin.RouterGroup) {
	app := api.App.SearchApi
	r.GET("search/article", middleware.BindQueryMiddleware[search_api.ArticleSearchRequest], app.ArticleSearchView)
	r.GET("search/text", middleware.BindQueryMiddleware[search_api.TextSearchRequest], app.TextSearchView)
}

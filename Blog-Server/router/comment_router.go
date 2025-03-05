// router/site_router.go
package router

import (
	"Blog-Server/api"
	"Blog-Server/api/comment_api"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.RouterGroup) {
	app := api.App.CommentApi
	r.POST("comment", middleware.AuthMiddleware, middleware.BindJsonMiddleware[comment_api.CommentCreateRequest], app.CommentCreateView)
	r.GET("comment/tree/:id", middleware.BindUriMiddleware[models.IDRequest], app.CommentTreeView)
	r.GET("comment", middleware.AuthMiddleware, middleware.BindQueryMiddleware[comment_api.CommentListRequest], app.CommentListView)
}

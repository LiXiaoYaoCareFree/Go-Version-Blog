package router

import (
	"Blog-Server/api"
	"Blog-Server/api/chat_api"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func ChatRouter(r *gin.RouterGroup) {
	app := api.App.ChatApi
	r.GET("chat", middleware.AuthMiddleware, middleware.BindQueryMiddleware[chat_api.ChatListRequest], app.ChatListView)
}

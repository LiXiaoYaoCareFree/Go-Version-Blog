package router

import (
	"Blog-Server/api"
	"Blog-Server/api/chat_api"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"github.com/gin-gonic/gin"
)

func ChatRouter(r *gin.RouterGroup) {
	app := api.App.ChatApi
	r.GET("chat", middleware.AuthMiddleware, middleware.BindQueryMiddleware[chat_api.ChatListRequest], app.ChatListView)
	r.GET("chat/session", middleware.AuthMiddleware, middleware.BindQueryMiddleware[chat_api.SessionListRequest], app.SessionListView)
	r.DELETE("chat", middleware.AuthMiddleware, middleware.BindJsonMiddleware[models.RemoveRequest], app.UserChatDeleteView)
	r.DELETE("chat/user/:id", middleware.AuthMiddleware, middleware.BindUriMiddleware[models.IDRequest], app.UserChatDeleteByUserView)
	r.POST("chat/read/:id", middleware.AuthMiddleware, middleware.BindUriMiddleware[models.IDRequest], app.ChatReadView)
	r.GET("chat/ws", app.ChatView)
}

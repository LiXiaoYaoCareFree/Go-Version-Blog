package api

import (
	"Blog-Server/api/ai_api"
	"Blog-Server/api/article_api"
	"Blog-Server/api/banner_api"
	"Blog-Server/api/captcha_api"
	"Blog-Server/api/chat_api"
	"Blog-Server/api/comment_api"
	"Blog-Server/api/focus_api"
	"Blog-Server/api/global_notification_api"
	"Blog-Server/api/image_api"
	"Blog-Server/api/log_api"
	"Blog-Server/api/search_api"
	"Blog-Server/api/site_api"
	"Blog-Server/api/site_msg_api"
	"Blog-Server/api/user_api"
)

type Api struct {
	SiteApi               site_api.SiteApi
	LogApi                log_api.LogApi
	ImageApi              image_api.ImageApi
	BannerApi             banner_api.BannerApi
	CaptchaApi            captcha_api.CaptchaApi
	UserApi               user_api.UserApi
	ArticleApi            article_api.ArticleApi
	CommentApi            comment_api.CommentApi
	SiteMsgApi            site_msg_api.SiteMsgApi
	GlobalNotificationApi global_notification_api.GlobalNotificationApi
	FocusApi              focus_api.FocusApi
	ChatApi               chat_api.ChatApi
	SearchApi             search_api.SearchApi
	AiApi                 ai_api.AiApi
}

var App = Api{}

package api

import (
	"Blog-Server/api/article_api"
	"Blog-Server/api/banner_api"
	"Blog-Server/api/captcha_api"
	"Blog-Server/api/image_api"
	"Blog-Server/api/log_api"
	"Blog-Server/api/site_api"
	"Blog-Server/api/user_api"
)

type Api struct {
	SiteApi    site_api.SiteApi
	LogApi     log_api.LogApi
	ImageApi   image_api.ImageApi
	BannerApi  banner_api.BannerApi
	CaptchaApi captcha_api.CaptchaApi
	UserApi    user_api.UserApi
	ArticleApi article_api.ArticleApi
}

var App = Api{}

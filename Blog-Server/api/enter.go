package api

import (
	"Blog-Server/api/banner_api"
	"Blog-Server/api/captcha_api"
	"Blog-Server/api/image_api"
	"Blog-Server/api/log_api"
	"Blog-Server/api/site_api"
)

type Api struct {
	SiteApi    site_api.SiteApi
	LogApi     log_api.LogApi
	ImageApi   image_api.ImageApi
	BannerApi  banner_api.BannerApi
	CaptchaApi captcha_api.CaptchaApi
}

var App = Api{}

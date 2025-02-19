package api

import (
	"Blog-Server/api/log_api"
	"Blog-Server/api/site_api"
)

type Api struct {
	SiteApi site_api.SiteApi
	LogApi  log_api.LogApi
}

var App = Api{}

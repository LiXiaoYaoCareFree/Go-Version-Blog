package api

import "Blog-Server/api/site_api"

type Api struct {
	SiteApi site_api.SiteApi
}

var App = Api{}

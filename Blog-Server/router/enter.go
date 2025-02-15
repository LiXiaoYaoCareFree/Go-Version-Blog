package router

import (
	"Blog-Server/global"
	"Blog-Server/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	nr := r.Group("/api")

	nr.Use(middleware.LogMiddleware)
	SiteRouter(nr)

	addr := global.Config.System.Addr()
	r.Run(addr)
}

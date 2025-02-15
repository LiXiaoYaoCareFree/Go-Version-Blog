package router

import (
	"Blog-Server/global"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	nr := r.Group("/api")

	SiteRouter(nr)

	addr := global.Config.System.Addr()
	r.Run(addr)
}

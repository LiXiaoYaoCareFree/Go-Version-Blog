// api/site_api/enter.go
package site_api

import (
	"Blog-Server/common/res"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	res.OkWithData("xx", c)
	return
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required" label:"年龄"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OkWithMsg("更新成功", c)
	return
}

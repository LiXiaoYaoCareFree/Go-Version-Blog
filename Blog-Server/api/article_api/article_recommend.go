package article_api

import (
	"Blog-Server/common"
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"github.com/gin-gonic/gin"
)

type ArticleRecommendResponse struct {
	ID        uint   `json:"id" gorm:"column:id"`
	Title     string `json:"title" gorm:"column:title"`
	LookCount int    `json:"lookCount" gorm:"column:lookCount"`
}

func (ArticleApi) ArticleRecommendView(c *gin.Context) {
	cr := middleware.GetBind[common.PageInfo](c)

	var list = make([]ArticleRecommendResponse, 0)
	global.DB.Model(models.ArticleModel{}).
		Order("look_count desc").
		Where("date(created_at) = date(now())").
		Limit(cr.Limit).Select("id", "title", "look_count").Scan(&list)

	res.OkWithList(list, len(list), c)
}

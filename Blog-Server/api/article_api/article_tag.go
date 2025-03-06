package article_api

import (
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/models"
	"Blog-Server/models/ctype"
	"Blog-Server/models/enum"
	"Blog-Server/utils"
	"Blog-Server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleTagOptionsView(c *gin.Context) {
	claims := jwts.GetClaims(c)

	var articleList []models.ArticleModel
	global.DB.Find(&articleList, "user_id = ? and status = ?", claims.UserID, enum.ArticleStatusPublished)

	var tagList ctype.List
	for _, model := range articleList {
		tagList = append(tagList, model.TagList...)
	}
	tagList = utils.Unique(tagList)
	var list = make([]models.OptionsResponse[string], 0)
	for _, s := range tagList {
		list = append(list, models.OptionsResponse[string]{
			Label: s,
			Value: s,
		})
	}
	res.OkWithData(list, c)
}

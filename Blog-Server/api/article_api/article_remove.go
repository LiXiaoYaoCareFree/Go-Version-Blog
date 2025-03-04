package article_api

import (
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	cr := middleware.GetBind[models.RemoveRequest](c)

	var list []models.ArticleModel
	global.DB.Find(&list, "id in ?", cr.IDList)

	if len(list) > 0 {
		err := global.DB.Delete(&list).Error
		if err != nil {
			res.FailWithMsg("删除失败", c)
			return
		}
	}

	res.OkWithMsg(fmt.Sprintf("删除成功 成功删除%d条", len(list)), c)
}

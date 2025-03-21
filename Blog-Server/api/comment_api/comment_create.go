package comment_api

import (
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"Blog-Server/models/enum"
	"Blog-Server/service/comment_service"
	"Blog-Server/service/message_service"
	"Blog-Server/service/redis_service/redis_article"
	"Blog-Server/service/redis_service/redis_comment"
	"Blog-Server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type CommentCreateRequest struct {
	Content   string `json:"content" binding:"required"`
	ArticleID uint   `json:"articleID" binding:"required"`
	ParentID  *uint  `json:"parentID"` // 父评论id
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	cr := middleware.GetBind[CommentCreateRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "id = ? and status = ?", cr.ArticleID, enum.ArticleStatusPublished).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	claims := jwts.GetClaims(c)

	model := models.CommentModel{
		Content:   cr.Content,
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
		ParentID:  cr.ParentID,
	}
	// 去找这个评论的根评论
	if cr.ParentID != nil {
		// 有父评论
		parentList := comment_service.GetParents(*cr.ParentID)
		// 判断父评论的层级是否满足
		if len(parentList)+1 > global.Config.Site.Article.CommentLine {
			res.FailWithMsg("评论层级达到限制", c)
			return
		}
		if len(parentList) > 0 {
			model.RootParentID = &parentList[len(parentList)-1].ID
			for _, commentModel := range parentList {
				redis_comment.SetCacheApply(commentModel.ID, 1)
			}
			// 给父评论的用有人发消息
			defer func() {
				go message_service.InsertCommentMessage(model)
			}()
		}
	}

	err = global.DB.Create(&model).Error
	if err != nil {
		res.FailWithMsg("发布评论失败", c)
		return
	}
	go message_service.InsertCommentMessage(model)
	redis_article.SetCacheComment(cr.ArticleID, 1)
	res.OkWithMsg("发布评论成功", c)

}

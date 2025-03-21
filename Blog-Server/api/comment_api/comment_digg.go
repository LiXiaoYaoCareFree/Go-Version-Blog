package comment_api

import (
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"Blog-Server/service/message_service"
	"Blog-Server/service/redis_service/redis_comment"
	"Blog-Server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentDiggView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)

	var comment models.CommentModel
	err := global.DB.Take(&comment, cr.ID).Error
	if err != nil {
		res.FailWithMsg("评论不存在", c)
		return
	}

	claims := jwts.GetClaims(c)

	// 查一下之前有没有点过
	var userDiggComment models.CommentDiggModel
	err = global.DB.Take(&userDiggComment, "user_id = ? and comment_id = ?", claims.UserID, comment.ID).Error
	if err != nil {
		// 点赞
		model := models.CommentDiggModel{
			UserID:    claims.UserID,
			CommentID: cr.ID,
		}
		err = global.DB.Create(&model).Error
		if err != nil {
			res.FailWithMsg("点赞失败", c)
			return
		}
		redis_comment.SetCacheDigg(cr.ID, 1)
		// 给这个评论的拥有人发消息
		message_service.InsertDiggCommentMessage(model)
		res.OkWithMsg("点赞成功", c)
		return
	}
	// 取消点赞
	global.DB.Delete(&userDiggComment)
	res.OkWithMsg("取消点赞成功", c)
	redis_comment.SetCacheDigg(cr.ID, -1)
	return
}

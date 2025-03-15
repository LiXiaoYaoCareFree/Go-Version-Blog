package comment_api

import (
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"Blog-Server/models/enum"
	"Blog-Server/models/enum/relationship_enum"
	"Blog-Server/service/comment_service"
	"Blog-Server/service/focus_service"
	"Blog-Server/utils"
	"Blog-Server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentTreeView(c *gin.Context) {
	var cr = middleware.GetBind[models.IDRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	var userRelationMap = map[uint]relationship_enum.Relation{}
	var userDiggCommentMap = map[uint]bool{}
	claims, err := jwts.ParseTokenByGin(c)
	if err == nil && claims != nil {
		// 登录了
		var commentList []models.CommentModel // 文章的评论id列表
		global.DB.Find(&commentList, "article_id = ?", cr.ID)

		if len(commentList) > 0 {
			// 查我点赞的评论id列表
			var commentIDList []uint
			var userIDList []uint
			for _, model := range commentList {
				commentIDList = append(commentIDList, model.ID)
				userIDList = append(userIDList, model.UserID)
			}
			userIDList = utils.Unique(userIDList) // 对用户id去重
			userRelationMap = focus_service.CalcUserPatchRelationship(claims.UserID, userIDList)
			var commentDiggList []models.CommentDiggModel
			global.DB.Find(&commentDiggList, "user_id = ? and comment_id in ?", claims.UserID, commentIDList)
			for _, model := range commentDiggList {
				userDiggCommentMap[model.CommentID] = true
			}
		}
	}
	// 把根评论查出来
	var commentList []models.CommentModel
	global.DB.Order("created_at desc").Find(&commentList, "article_id = ? and parent_id is null", cr.ID)
	var list = make([]comment_service.CommentResponse, 0)
	for _, model := range commentList {
		response := comment_service.GetCommentTreeV4(model.ID, userRelationMap, userDiggCommentMap)
		list = append(list, *response)
	}

	res.OkWithList(list, len(list), c)
}

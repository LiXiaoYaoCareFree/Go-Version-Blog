// api/user_api/user_base_info.go
package user_api

import (
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/models"
	"Blog-Server/models/enum/relationship_enum"
	"Blog-Server/service/focus_service"
	"Blog-Server/service/redis_service/redis_user"
	"Blog-Server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserBaseInfoResponse struct {
	UserID       uint                       `json:"userID"`
	CodeAge      int                        `json:"codeAge"`
	Avatar       string                     `json:"avatar"`
	Nickname     string                     `json:"nickname"`
	LookCount    int                        `json:"lookCount"`
	ArticleCount int                        `json:"articleCount"`
	FansCount    int                        `json:"fansCount"`
	FollowCount  int                        `json:"followCount"`
	Place        string                     `json:"place"`       // ip归属地
	OpenCollect  bool                       `json:"openCollect"` // 公开我的收藏
	OpenFollow   bool                       `json:"openFollow"`  // 公开我的关注
	OpenFans     bool                       `json:"openFans"`    // 公开我的粉丝
	HomeStyleID  uint                       `json:"homeStyleID"` // 主页样式的id
	Relation     relationship_enum.Relation `json:"relation"`    // 与登录人的关系
}

func (UserApi) UserBaseInfoView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var user models.UserModel
	err = global.DB.Preload("UserConfModel").Preload("ArticleList").Take(&user, cr.ID).Error
	if err != nil {
		res.FailWithMsg("不存在的用户", c)
		return
	}

	data := UserBaseInfoResponse{
		UserID:       user.ID,
		CodeAge:      user.CodeAge(),
		Avatar:       user.Avatar,
		Nickname:     user.Nickname,
		LookCount:    user.UserConfModel.LookCount + redis_user.GetCacheLook(cr.ID),
		ArticleCount: len(user.ArticleList),
		FansCount:    0,
		FollowCount:  0,
		Place:        user.Addr,
		OpenCollect:  user.UserConfModel.OpenCollect,
		OpenFollow:   user.UserConfModel.OpenFollow,
		OpenFans:     user.UserConfModel.OpenFans,
		HomeStyleID:  user.UserConfModel.HomeStyleID,
	}

	claims, err := jwts.ParseTokenByGin(c)
	if err == nil && claims != nil {
		data.Relation = focus_service.CalcUserRelationship(claims.UserID, cr.ID)
	}

	var focusList []models.UserFocusModel
	global.DB.Find(&focusList, "user_id = ? or focus_user_id = ?", cr.ID, cr.ID)
	for _, model := range focusList {
		if model.UserID == cr.ID {
			data.FansCount++
		}
		if model.FocusUserID == cr.ID {
			data.FollowCount++
		}
	}

	redis_user.SetCacheLook(cr.ID, true)

	res.OkWithData(data, c)
}

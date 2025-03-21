package focus_api

import (
	"Blog-Server/common"
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"Blog-Server/models/enum/relationship_enum"
	"Blog-Server/service/focus_service"
	"Blog-Server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type FocusApi struct {
}
type FocusUserRequest struct {
	FocusUserID uint `json:"focusUserID" binding:"required"`
}

// FocusUserView 登录人关注用户
func (FocusApi) FocusUserView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserRequest](c)

	claims := jwts.GetClaims(c)
	if cr.FocusUserID == claims.UserID {
		res.FailWithMsg("你时刻都在关注自己", c)
		return
	}
	// 查关注的用户是否存在
	var user models.UserModel
	err := global.DB.Take(&user, cr.FocusUserID).Error
	if err != nil {
		res.FailWithMsg("关注用户不存在", c)
		return
	}

	// 查之前是否已经关注过他了
	var focus models.UserFocusModel
	err = global.DB.Take(&focus, "user_id = ? and focus_user_id = ?", claims.UserID, user.ID).Error
	if err == nil {
		res.FailWithMsg("请勿重复关注", c)
		return
	}

	// 每天关注是不是应该有个限度？
	// 每天的取关也要有个限度？

	// 关注
	global.DB.Create(&models.UserFocusModel{
		UserID:      claims.UserID,
		FocusUserID: cr.FocusUserID,
	})

	res.OkWithMsg("关注成功", c)
	return
}

type FocusUserListRequest struct {
	common.PageInfo
	FocusUserID uint `form:"focusUserID"`
	UserID      uint `form:"userID"` // 查用户的关注
}
type UserListResponse struct {
	UserID       uint                       `json:"userID"`
	UserNickname string                     `json:"userNickname"`
	UserAvatar   string                     `json:"userAvatar"`
	UserAbstract string                     `json:"userAbstract"`
	Relationship relationship_enum.Relation `json:"relationship"`
	CreatedAt    time.Time                  `json:"createdAt"`
}

// FocusUserListView 我的关注和用户的关注
func (FocusApi) FocusUserListView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserListRequest](c)

	claims, err := jwts.ParseTokenByGin(c)

	if cr.UserID != 0 {
		// 传了用户id，我就查这个人关注的用户列表
		var userConf models.UserConfModel
		err1 := global.DB.Take(&userConf, "user_id = ?", cr.UserID).Error
		if err1 != nil {
			res.FailWithMsg("用户配置信息不存在", c)
			return
		}
		if !userConf.OpenFollow {
			res.FailWithMsg("此用户未公开我的关注", c)
			return
		}

		// 如果你没登录。我就不允许你查第二页
		if err != nil || claims == nil {
			if cr.Limit > 10 || cr.Page > 1 {
				res.FailWithMsg("未登录用户只能显示第一页", c)
				return
			}
		}

	} else {
		if err != nil || claims == nil {
			res.FailWithMsg("请登录", c)
			return
		}
		cr.UserID = claims.UserID
	}

	query := global.DB.Where("")
	if cr.Key != "" {
		// 模糊匹配用户
		var userIDList []uint
		global.DB.Model(&models.UserModel{}).
			Where("nickname like ?", fmt.Sprintf("%%%s%%", cr.Key)).
			Select("id").Scan(&userIDList)
		if len(userIDList) > 0 {
			query.Where("focus_user_id in ?", userIDList)
		}
	}

	_list, count, _ := common.ListQuery(models.UserFocusModel{
		FocusUserID: cr.FocusUserID,
		UserID:      cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Where:    query,
		Preloads: []string{"FocusUserModel"},
	})

	var m = map[uint]relationship_enum.Relation{}
	if err == nil && claims != nil {
		var userIDList []uint
		for _, i2 := range _list {
			userIDList = append(userIDList, i2.FocusUserID)
		}
		m = focus_service.CalcUserPatchRelationship(claims.UserID, userIDList)

	}

	var list = make([]UserListResponse, 0)
	for _, model := range _list {
		list = append(list, UserListResponse{
			UserID:       model.FocusUserID,
			UserNickname: model.FocusUserModel.Nickname,
			UserAvatar:   model.FocusUserModel.Avatar,
			UserAbstract: model.FocusUserModel.Abstract,
			Relationship: m[model.FocusUserID],
			CreatedAt:    model.CreatedAt,
		})
	}

	res.OkWithList(list, count, c)
}

// FansUserListView 我的粉丝和用户的粉丝
func (FocusApi) FansUserListView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserListRequest](c)
	claims, err := jwts.ParseTokenByGin(c)
	if cr.UserID != 0 {
		// 传了用户id，我就查这个人的粉丝列表
		var userConf models.UserConfModel
		err1 := global.DB.Take(&userConf, "user_id = ?", cr.UserID).Error
		if err1 != nil {
			res.FailWithMsg("用户配置信息不存在", c)
			return
		}
		if !userConf.OpenFans {
			res.FailWithMsg("此用户未公开我的粉丝", c)
			return
		}
		// 如果你没登录。我就不允许你查第二页
		if err != nil || claims == nil {
			if cr.Limit > 10 || cr.Page > 1 {
				res.FailWithMsg("未登录用户只能显示第一页", c)
				return
			}
		}
	} else {
		if err != nil || claims == nil {
			res.FailWithMsg("请登录", c)
			return
		}
		cr.UserID = claims.UserID
	}

	query := global.DB.Where("")
	if cr.Key != "" {
		// 模糊匹配用户
		var userIDList []uint
		global.DB.Model(&models.UserModel{}).
			Where("nickname like ?", fmt.Sprintf("%%%s%%", cr.Key)).
			Select("id").Scan(&userIDList)
		if len(userIDList) > 0 {
			query.Where("user_id in ?", userIDList)
		}
	}

	_list, count, _ := common.ListQuery(models.UserFocusModel{
		FocusUserID: cr.UserID,
		UserID:      cr.FocusUserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Where:    query,
		Preloads: []string{"UserModel"},
	})
	var m = map[uint]relationship_enum.Relation{}
	if err == nil && claims != nil {
		var userIDList []uint
		for _, i2 := range _list {
			userIDList = append(userIDList, i2.UserID)
		}
		m = focus_service.CalcUserPatchRelationship(claims.UserID, userIDList)
	}

	var list = make([]UserListResponse, 0)
	for _, model := range _list {
		list = append(list, UserListResponse{
			UserID:       model.UserID,
			UserNickname: model.UserModel.Nickname,
			UserAvatar:   model.UserModel.Avatar,
			UserAbstract: model.UserModel.Abstract,
			CreatedAt:    model.CreatedAt,
			Relationship: m[model.UserID],
		})
	}

	res.OkWithList(list, count, c)
}

// UnFocusUserView 登录人取关用户
func (FocusApi) UnFocusUserView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserRequest](c)

	claims := jwts.GetClaims(c)
	if cr.FocusUserID == claims.UserID {
		res.FailWithMsg("你无法取关自己", c)
		return
	}
	// 查关注的用户是否存在
	var user models.UserModel
	err := global.DB.Take(&user, cr.FocusUserID).Error
	if err != nil {
		res.FailWithMsg("取关用户不存在", c)
		return
	}

	// 查之前是否已经关注过他了
	var focus models.UserFocusModel
	err = global.DB.Take(&focus, "user_id = ? and focus_user_id = ?", claims.UserID, user.ID).Error
	if err != nil {
		res.FailWithMsg("未关注此用户", c)
		return
	}
	// 每天的取关也要有个限度？
	// 取关
	global.DB.Delete(&focus)
	res.OkWithMsg("取消关注成功", c)
	return
}

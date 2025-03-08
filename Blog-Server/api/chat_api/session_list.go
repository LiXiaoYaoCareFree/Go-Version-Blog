package chat_api

import (
	"Blog-Server/common"
	"Blog-Server/common/res"
	"Blog-Server/global"
	"Blog-Server/middleware"
	"Blog-Server/models"
	"Blog-Server/utils/jwts"
	"Blog-Server/utils/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SessionListRequest struct {
	common.PageInfo
}

type SessionTable struct {
	SU        uint   `gorm:"column:sU"`
	RU        uint   `gorm:"column:rU"`
	MaxDate   string `gorm:"column:maxDate"`
	Count     int    `gorm:"column:c"`
	NewChatID uint   `gorm:"column:newChatID"`
}

func (ChatApi) SessionListView(c *gin.Context) {

	cr := middleware.GetBind[SessionListRequest](c)
	claims := jwts.GetClaims(c)
	// 查我删了哪些消息
	var deletedIDList []uint
	global.DB.Model(models.UserChatActionModel{}).
		Where("user_id = ? and is_delete = ?", claims.UserID, true).
		Select("chat_id").Scan(&deletedIDList)

	query := global.DB.Where("")

	var column = fmt.Sprintf("(select id from chat_models where ((send_user_id = sU and rev_user_id = rU) or (send_user_id = rU and rev_user_id = sU))  order by created_at desc limit 1) as newChatID")
	if len(deletedIDList) > 0 {
		query.Where("id not in ?", deletedIDList)
		column = fmt.Sprintf("(select id from chat_models where ((send_user_id = sU and rev_user_id = rU) or (send_user_id = rU and rev_user_id = sU)) and id not in %s order by created_at desc limit 1) as newChatID", sql.ConvertSliceSql(deletedIDList))
	}
	var _list []SessionTable
	global.DB.Model(models.ChatModel{}).
		Select(
			"least(send_user_id, rev_user_id)    as sU",
			"greatest(send_user_id, rev_user_id) as rU",
			"max(created_at)       as maxDate",
			"count(*)         as c",
			column,
		).
		Where(query).
		Where("(send_user_id = ? or rev_user_id = ?)", claims.UserID, claims.UserID).
		Group("least(send_user_id, rev_user_id)").
		Group("greatest(send_user_id, rev_user_id)").
		Order("maxDate desc").
		Limit(cr.GetLimit()).Offset(cr.GetOffset()).Scan(&_list)

	var count int
	global.DB.Select("count(*)").Table("(?) as x",
		global.DB.
			Model(models.ChatModel{}).
			Select("count(*)").
			Where("(send_user_id = ? or rev_user_id = ?)", claims.UserID, claims.UserID).
			Group("least(send_user_id, rev_user_id)").
			Group("greatest(send_user_id, rev_user_id)"),
	).Scan(&count)

	res.OkWithList(_list, count, c)
}

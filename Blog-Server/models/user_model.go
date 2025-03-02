// models/user_model.go
package models

import (
	"Blog-Server/models/enum"
	"time"
)

type UserModel struct {
	Model
	Username       string                  `gorm:"size:32" json:"username"`
	Nickname       string                  `gorm:"size:32" json:"nickname"`
	Avatar         string                  `gorm:"size:256" json:"avatar"`
	Abstract       string                  `gorm:"size:256" json:"abstract"`
	RegisterSource enum.RegisterSourceType `json:"registerSource"`
	Password       string                  `gorm:"size:64" json:"-"`
	Email          string                  `gorm:"size:256" json:"email"`
	OpenID         string                  `gorm:"size:64" json:"openID"` // 第三方登陆的唯一id
	Role           enum.RoleType           `json:"role"`                  // 角色  1  管理员  2  普通用户 3  访客
	UserConfModel  *UserConfModel          `gorm:"foreignKey:UserID"  json:"-"`
}

type UserConfModel struct {
	UserID             uint       `gorm:"unique" json:"userID"`
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"`
	UpdateUsernameDate *time.Time `json:"updateUsernameDate"` // 上次修改用户名的时间
	OpenCollect        bool       `json:"openCollect"`        // 公开我的收藏
	OpenFans           bool       `json:"openFans"`           // 公开我的粉丝
	OpenFollow         bool       `json:"openFollow"`         // 公开我的关注
	HomeStyleID        uint       `json:"homeStyleID"`        //主页样式
}

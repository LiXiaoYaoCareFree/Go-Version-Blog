package models

import (
	"Blog-Server/models/ctype/chat_msg"
	"Blog-Server/models/enum/chat_msg_type"
)

type ChatModel struct {
	Model
	SendUserID    uint                  `json:"sendUserID"`
	SendUserModel UserModel             `gorm:"foreignKey:SendUserID"  json:"-"`
	RevUserID     uint                  `json:"revUserID"`
	RevUserModel  UserModel             `gorm:"foreignKey:RevUserID"  json:"-"`
	MsgType       chat_msg_type.MsgType `json:"msgType"` // 消息类型
	Msg           chat_msg.ChatMsg      `gorm:"type:longtext;serializer:json" json:"msg"`
}

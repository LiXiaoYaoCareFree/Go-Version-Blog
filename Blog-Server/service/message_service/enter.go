package message_service

import (
	"Blog-Server/global"
	"Blog-Server/models"
	"Blog-Server/models/enum/message_type_enum"
	"github.com/sirupsen/logrus"
)

// InsertCommentMessage 插入一条评论消息
func InsertCommentMessage(model models.CommentModel) {
	global.DB.Preload("UserModel").Preload("ArticleModel").Take(&model)
	err := global.DB.Create(&models.MessageModel{
		Type:               message_type_enum.CommentType,
		RevUserID:          model.ArticleModel.UserID,
		ActionUserID:       model.UserID,
		ActionUserNickname: model.UserModel.Nickname,
		ActionUserAvatar:   model.UserModel.Avatar,
		Content:            model.Content,
		ArticleID:          model.ArticleID,
		ArticleTitle:       model.ArticleModel.Title,
		CommentID:          model.ID,
	}).Error
	if err != nil {
		logrus.Error(err)
	}

}

// InsertApplyMessage 插入一条评论回复消息
func InsertApplyMessage(model models.CommentModel) {
	// TODO：回复评论的人和自己是同一个人，要不要通知？
	global.DB.Preload("ParentModel").Preload("UserModel").Preload("ArticleModel").Take(&model)
	err := global.DB.Create(&models.MessageModel{
		Type:               message_type_enum.ApplyType,
		RevUserID:          model.ParentModel.UserID,
		ActionUserID:       model.UserID,
		ActionUserNickname: model.UserModel.Nickname,
		ActionUserAvatar:   model.UserModel.Avatar,
		Content:            model.Content,
		ArticleID:          model.ArticleID,
		ArticleTitle:       model.ArticleModel.Title,
		CommentID:          model.ID,
	}).Error
	if err != nil {
		logrus.Error(err)
	}

}

package flags

import (
	"Blog-Server/global"
	"Blog-Server/models"
	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},
		&models.UserConfModel{},
		&models.ArticleModel{},
		&models.CategoryModel{},
		&models.ArticleDiggModel{},
		&models.CollectModel{},
		&models.UserArticleCollectModel{},
		&models.ImageModel{},
		&models.UserTopArticleModel{},
		&models.UserArticleLookHistoryModel{},
		&models.CommentModel{},
		&models.BannerModel{},
		&models.LogModel{},
		&models.GlobalNotificationModel{},
		&models.ImageModel{},
		&models.UserLoginModel{},              // 用户登陆记录表
		&models.CommentDiggModel{},            // 用户点赞评论表
		&models.MessageModel{},                // 站内信表
		&models.UserMessageConfModel{},        // 用户消息配置表
		&models.UserGlobalNotificationModel{}, // 用户全局消息表
		&models.UserFocusModel{},              // 好友关系表
		&models.ChatModel{},                   // 对话表
		&models.UserChatActionModel{},         // 用户操作对话表，读取，删除
		&models.TextModel{},                   // 全文搜索表
		&models.SiteFlowModel{},               // 网站流量表
	)
	if err != nil {
		logrus.Errorf("数据库迁移失败 %s", err)
		return
	}
	logrus.Infof("数据库迁移成功！")
}

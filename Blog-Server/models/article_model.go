package models

import _ "embed"

type ArticleModel struct {
	Model
	Title        string    `gorm:"size:32" json:"title"`
	Abstract     string    `gorm:"size:256" json:"abstract"`
	Content      string    `json:"content"`
	CategoryID   uint      `json:"categoryID"`
	TagList      []string  `gorm:"type:longtext;serializer:json" json:"tagList"` // 标签列表
	Cover        string    `gorm:"256" json:"cover"`
	UserID       uint      `json:"UserID"`
	UserModel    UserModel `gorm:"foreignKey:UserID" json:"-"`
	LookCount    int       `json:"lookCount"`
	DiggCount    int       `json:"diggCount"`
	CommentCount int       `json:"commentCount"`
	CollectCount int       `json:"collectCount"`
	OpenComment  bool      `json:"openComment"`
	Status       int8      `json:"status"` //状态  草稿 审核中 已发布
}

//go:embed mappings/article_mapping.json
var articleMapping string

func (ArticleModel) Mapping() string {
	return articleMapping
}

func (ArticleModel) Index() string {
	return "article_index"
}

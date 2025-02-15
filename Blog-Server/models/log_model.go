package models

type LogModel struct {
	LogType   int8      `json:"logType"` // 日志类型
	Title     string    `gorm:"size:64" json:"title"`
	Content   string    `json:"content"`
	Level     int8      `json:"level"`  // 日志级别
	UserID    uint      `json:"userID"` // 用户id
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:32" json:"ip"`
	Addr      string    `gorm:"size:64" json:"addr"`
	IsRead    bool      `json:"isRead"` // 是否已读
}

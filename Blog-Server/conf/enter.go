package conf

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	Jwt    Jwt    `yaml:"jwt"`
	DB     []DB   `yaml:"db"` // 数据库连接列表
	Redis  Redis  `yaml:"redis"`
	Site   Site   `yaml:"site"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	QiNiu  QiNiu  `yaml:"qiNiu"`
	Ai     Ai     `yaml:"ai"`
	Upload Upload `yaml:"upload"`
	ES     ES     `yaml:"es"`
}

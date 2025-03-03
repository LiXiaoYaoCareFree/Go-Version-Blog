// core/init_mysql_es.go
package core

import (
	"Blog-Server/global"
	river "Blog-Server/service/river_service"
	"github.com/sirupsen/logrus"
)

func InitMysqlES() {
	if !global.Config.River.Enable {
		logrus.Infof("关闭mysql同步操作")
		return
	}
	r, err := river.NewRiver()
	if err != nil {
		logrus.Fatal(err)
	}
	go r.Run()
}

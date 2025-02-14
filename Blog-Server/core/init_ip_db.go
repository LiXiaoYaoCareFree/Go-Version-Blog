package core

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
)

var searcher *xdb.Searcher

func InitIPDB() {
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		logrus.Fatalf("ip地址数据库加载失败 %s", err)
		return
	}

	searcher = _searcher
}

func GetIpAddr(ip string) (addr string, err error) {
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		return
	}
	return region, nil
}

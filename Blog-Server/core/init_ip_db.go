package core

import (
	ipUtils "Blog-Server/utils/ip"
	_ "embed"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
	"strings"
)

var searcher *xdb.Searcher

//go:embed ip2region.xdb
var addrDB []byte

func InitIPDB() {
	_searcher, err := xdb.NewWithBuffer(addrDB)
	if err != nil {
		logrus.Fatalf("ip地址数据库加载失败 %s", err)
		return
	}
	searcher = _searcher
}

func GetIpAddr(ip string) (addr string) {
	if ipUtils.HasLocalIPAddr(ip) {
		return "内网"
	}

	region, err := searcher.SearchByStr(ip)
	if err != nil {
		logrus.Warnf("错误的ip地址 %s", err)
		return "异常地址"
	}

	_addrList := strings.Split(region, "|")
	if len(_addrList) != 5 {
		logrus.Fatalf("异常的ip地址 %s", ip)
		return "未知地址"
	}

	// _addrList 五个部分
	// 国家	0	省份	   市     运营商
	country := _addrList[0]
	province := _addrList[2]
	city := _addrList[3]

	if province != "0" && city != "0" {
		return fmt.Sprintf("%s·%s", province, city)
	}

	if country != "0" && province != "0" {
		return fmt.Sprintf("%s·%s", country, province)
	}
	if country != "0" {
		return country
	}

	return region
}

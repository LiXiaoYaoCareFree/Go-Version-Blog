package main

import (
	"Blog-Server/core"
	"Blog-Server/flags"
	"Blog-Server/global"
	"Blog-Server/models"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	//err := global.DB.Create(&models.ArticleModel{
	//	Title:   "嘻嘻嘻",
	//	TagList: ctype.List{"python", "go"},
	//}).Error
	//fmt.Println(err)
	var list1 []models.ArticleModel
	global.DB.Find(&list1)
	fmt.Println(list1)
}

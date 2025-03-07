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

	fmt.Println(models.CalcUserRelationship(1, 2))
	fmt.Println(models.CalcUserRelationship(1, 3))
	fmt.Println(models.CalcUserRelationship(4, 1))
	fmt.Println(models.CalcUserRelationship(5, 1))
	fmt.Println(models.CalcUserRelationship(1, 4))

}

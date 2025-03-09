package main

import (
	"Blog-Server/models"
	"Blog-Server/service/text_service"
	"fmt"
	"os"
)

func main() {
	byteData, _ := os.ReadFile("text.md")
	list := text_service.MdContentTransformation(models.ArticleModel{
		Model:   models.Model{ID: 4},
		Title:   "xxx",
		Content: string(byteData),
	})

	fmt.Println(list)

}

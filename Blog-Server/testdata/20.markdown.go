package main

import (
	"Blog-Server/utils/markdown"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var md = `

# 这是一级标题

> 这是应用

![xxx](jjdfnghjdf)

`

func main() {
	html := markdown.MdToHtml(md)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	if err != nil {
		fmt.Println(err)
		return
	}
	htmlText := doc.Text()
	fmt.Println(htmlText)
}

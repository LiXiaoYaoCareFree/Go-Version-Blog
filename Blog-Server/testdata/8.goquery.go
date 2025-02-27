package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func main() {
	reader, err := os.Open("uploads/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	selection := doc.Find("title")
	doc.Find("head").AppendHtml("<meta name=\"keywords1\" content=\"枫枫知道,网站,开发,程序员,golang\">")
	//fmt.Println(selection.Text())
	selection.SetText("CodeForge")
	selection.SetAttr("", "")
	html, err := doc.Html()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(html)
}

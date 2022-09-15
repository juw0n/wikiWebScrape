package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getHTML(url string) *http.Response {
	resp, err := http.Get(url)
	checkErr(err)

	if resp.StatusCode > 400 {
		fmt.Println("status code", resp.StatusCode)
	}

	return resp
}

func getLinks(doc *goquery.Document) {
	doc.Find("div.mw-body-content").Each(func(i int, item *goquery.Selection) {
		a := item.Find("")
	})
}

func main() {
	// getting the page url
	url := "https://en.wikipedia.org/wiki/Web_scraping"

	resp := getHTML(url)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	getLinks(doc)
}

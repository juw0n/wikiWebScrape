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

func link(doc *goquery.Document) {
	doc.Find(".mw-parser-output")
}

func main() {
	url := "https://en.wikipedia.org/wiki/Web_scraping"

	resp := getHTML(url)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	link(doc)
}

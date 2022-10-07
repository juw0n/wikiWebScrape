package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// getting the page url
	url := "https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2334524.m570.l1313&_nkw=samsung&_sacat=0&LH_TitleDesc=0&_odkw=phone&_osacat=0"

	resp := getHTML(url)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	scrapeDate(doc)
}

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

func scrapeDate(doc *goquery.Document) {
	doc.Find("ul.srp-results>li.s-item").Each(func(i int, item *goquery.Selection) {
		a := item.Find("a.s-item__link")

		title := strings.TrimSpace(a.Text())
		url, _ := a.Attr("href")

		price_span := strings.TrimSpace(item.Find("span.s-item__price").Text())
		price := strings.Trim(price_span, " py6.")

		scrapedData := []string{title, price, url}

		writeCsv(scrapedData)
	})
}

func writeCsv(scrapedData []string) {
	filename := "data.csv"

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	checkErr(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(scrapedData)
	checkErr(err)
}
